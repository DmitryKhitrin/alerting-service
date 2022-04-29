package server

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func request(t *testing.T, ts *httptest.Server, method string, query string, body *bytes.Buffer) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+query, body)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	respBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestMetricServer(t *testing.T) {
	type want struct {
		code int
		body []string
	}
	// Mock App in future
	app := NewApp()
	r := getRouter(app)

	server := httptest.NewServer(r)
	defer server.Close()

	tests := []struct {
		description string
		requestURL  string
		method      string
		body        []byte
		expected    want
	}{
		{
			description: "200 Success Gauge",
			requestURL:  "/update/gauge/numberMetric/100",
			method:      http.MethodPost,
			body:        []byte(``),
			expected:    want{code: 200},
		},
		{
			description: "200 Success Gauge JSON",
			requestURL:  "/update",
			method:      http.MethodPost,
			body:        []byte(`{"id":"Alloc","type":"gauge","value":3459}`),
			expected:    want{code: 200},
		},
		{
			description: "200 Success Counter JSON",
			requestURL:  "/update",
			method:      http.MethodPost,
			body:        []byte(`{"id":"Alloc","type":"counter","delta":3459}`),
			expected:    want{code: 200},
		},
		{
			description: "400 Bad Counter JSON",
			requestURL:  "/update",
			method:      http.MethodPost,
			body:        []byte(`{"id":"Alloc","type":"counter","value":3459}`),
			expected:    want{code: 400},
		},
		{
			description: "400 Bad Gauge JSON",
			requestURL:  "/update",
			method:      http.MethodPost,
			body:        []byte(`{"id":"Alloc","type":"gauge","delta":3459}`),
			expected:    want{code: 400},
		},
		{
			description: "200 Success Counter",
			requestURL:  "/update/counter/PollCount/1",
			method:      http.MethodPost,
			body:        []byte(``),
			expected:    want{code: 200},
		},
		{
			description: "200 Get Counter",
			requestURL:  "/value/counter/PollCount",
			method:      http.MethodGet,
			body:        []byte(``),
			expected: want{
				code: 200,
				body: []string{"1"},
			},
		},
		{
			description: "200 Success Counter",
			requestURL:  "/update/counter/PollCount/10",
			method:      http.MethodPost,
			body:        []byte(``),
			expected:    want{code: 200},
		},
		{
			description: "200 Get Counter plus ten",
			requestURL:  "/value/counter/PollCount",
			method:      http.MethodGet,
			body:        []byte(``),
			expected: want{
				code: 200,
				body: []string{"11"},
			},
		},
		{
			description: "400 Parse string Gauge",
			requestURL:  "/update/gauge/MetricName/string",
			method:      http.MethodPost,
			body:        []byte(``),
			expected:    want{code: 400},
		},
		{
			description: "400 Parse string Counter",
			requestURL:  "/update/counter/MetricName/665g6",
			method:      http.MethodPost,
			body:        []byte(``),
			expected:    want{code: 400},
		},
		{
			description: "501 No such metric",
			requestURL:  "/update/type/MetricName/123",
			method:      http.MethodPost,
			body:        []byte(``),
			expected:    want{code: 501},
		},
		{
			description: "Get unknown Gauge",
			method:      http.MethodGet,
			body:        []byte(``),
			requestURL:  "/value/gauge/unknown",
			expected:    want{code: 404},
		},
		{
			description: "Get unknown Counter",
			method:      http.MethodGet,
			body:        []byte(``),
			requestURL:  "/value/counter/unknown",
			expected:    want{code: 404},
		},
		{
			description: "400 Get unknown type",
			method:      http.MethodGet,
			body:        []byte(``),
			requestURL:  "/value/type/name",
			expected:    want{code: 400},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			resp, body := request(t, server, tt.method, tt.requestURL, bytes.NewBuffer(tt.body))
			defer resp.Body.Close()
			assert.Equal(t, tt.expected.code, resp.StatusCode)
			for _, s := range tt.expected.body {
				assert.Equal(t, body, s)
			}
			assert.Equal(t, tt.expected.code, resp.StatusCode)
		})
	}
}
