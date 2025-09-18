package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		input http.Header
		want  string
	}{
		{input: "Apikey first", want: "first"},
		{input: "Apikey sedong", want: "second"},
	}

	for i, tc := range tests {
		got, err := GetAPIKey(tc.input)

	}

}
