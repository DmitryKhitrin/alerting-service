package utils

import (
	"reflect"
	"testing"
)

func TestParseURL(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    MetricParams
		wantErr bool
	}{
		{name: "Valid url", args: args{path: "/update/counter/Name/value"}, want: MetricParams{
			Name:  "Name",
			Value: "value",
		}, wantErr: false},
		{name: "Invalid url", args: args{path: "/update/counter/Name"}, want: MetricParams{
			Name:  "Name",
			Value: "",
		}, wantErr: true},
		{name: "Invalid url", args: args{path: "/update/counter/"}, want: MetricParams{
			Name:  "",
			Value: "",
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURL(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
