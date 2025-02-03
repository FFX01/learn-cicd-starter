package auth

import (
	"net/http"
	"testing"
)

func TestGetAuthKey(t *testing.T) {
	var tests = []struct {
		Header   http.Header
		Expected string
		Error    bool
	}{
		{
			Header:   http.Header{"Authorization": {"ApiKey 12345"}},
			Expected: "12345",
			Error:    false,
		},
		{
			Header:   http.Header{"Authorization": {"ApiKey"}},
			Expected: "",
			Error:    true,
		},
		{
			Header:   http.Header{"foobar": {"baz"}},
			Expected: "",
			Error:    true,
		},
		{
			Header:   http.Header{"": {""}},
			Expected: "",
			Error:    true,
		},
		{
			Header:   http.Header{"Authorization": {"bearer 12345"}},
			Expected: "",
			Error:    true,
		},
		{
			Header:   http.Header{},
			Expected: "",
			Error:    true,
		},
	}

	for _, tt := range tests {
		key, err := GetAPIKey(tt.Header)
		if tt.Error && err == nil {
			t.Error("Expected error, but got none", "key", key, "error", err)
		}
		if key != tt.Expected {
			t.Errorf("expected %s but got %s", tt.Expected, key)
		}
	}
}
