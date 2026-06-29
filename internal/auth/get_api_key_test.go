package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	apiKeyHeader := http.Header{}
	apiKeyHeader.Set("Authorization", "ApiKey 123456")

	bearerKeyHeader := http.Header{}
	bearerKeyHeader.Set("Authorization", "Bearer 123456")

	tests := map[string]struct {
		header   http.Header
		wantsErr bool
		wantsRes string
	}{
		"Has API Key":   {header: apiKeyHeader, wantsErr: false, wantsRes: "123456"},
		"No API Key":    {header: http.Header{}, wantsErr: false, wantsRes: ""},
		"Wrong API Key": {header: bearerKeyHeader, wantsErr: true, wantsRes: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.header)

			if err != nil && !tc.wantsErr {
				t.Fatalf("Got error when was not expected: %v", err.Error())
			}

			if got != tc.wantsRes {
				t.Fatalf("expected: %v, got: %v", tc.wantsRes, got)
			}
		})
	}
}
