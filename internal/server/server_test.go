package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func request(t *testing.T, ts *httptest.Server, method string, query string, body string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+query, strings.NewReader(body))
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
	r := getRouter()
	server := httptest.NewServer(r)
	defer server.Close()

	tests := []struct {
		description string
		requestURL  string
		method      string
		expected    want
	}{
		{
			description: "200 Success Gauge",
			requestURL:  "/update/gauge/numberMetric/100",
			method:      http.MethodPost,
			expected:    want{code: 200},
		},
		{
			description: "200 Success Counter",
			requestURL:  "/update/counter/PollCount/1",
			method:      http.MethodPost,
			expected:    want{code: 200},
		},
		{
			description: "200 Get Counter",
			requestURL:  "/value/counter/PollCount",
			method:      http.MethodGet,
			expected: want{
				code: 200,
				body: []string{"1"},
			},
		},
		{
			description: "200 Success Counter",
			requestURL:  "/update/counter/PollCount/10",
			method:      http.MethodPost,
			expected:    want{code: 200},
		},
		{
			description: "200 Get Counter plus ten",
			requestURL:  "/value/counter/PollCount",
			method:      http.MethodGet,
			expected: want{
				code: 200,
				body: []string{"11"},
			},
		},
		{
			description: "400 Parse string Gauge",
			requestURL:  "/update/gauge/MetricName/string",
			method:      http.MethodPost,
			expected:    want{code: 400},
		},
		{
			description: "400 Parse string Counter",
			requestURL:  "/update/counter/MetricName/665g6",
			method:      http.MethodPost,
			expected:    want{code: 400},
		},
		{
			description: "501 No such metric",
			requestURL:  "/update/type/MetricName/123",
			method:      http.MethodPost,
			expected:    want{code: 501},
		},
		{
			description: "Get unknown Gauge",
			method:      http.MethodGet,
			requestURL:  "/value/gauge/unknown",
			expected:    want{code: 404},
		},
		{
			description: "Get unknown Counter",
			method:      http.MethodGet,
			requestURL:  "/value/counter/unknown",
			expected:    want{code: 404},
		},
		{
			description: "400 Get unknown type",
			method:      http.MethodGet,
			requestURL:  "/value/type/name",
			expected:    want{code: 400},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			resp, body := request(t, server, tt.method, tt.requestURL, "")
			defer resp.Body.Close()
			assert.Equal(t, tt.expected.code, resp.StatusCode)
			for _, s := range tt.expected.body {
				assert.Equal(t, body, s)
			}
			assert.Equal(t, tt.expected.code, resp.StatusCode)
		})
	}
}
