package handlers

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeAndVerifyOriginalURL(t *testing.T) {
	tests := []struct {
		name          string
		body          string
		expectedURL   string
		expectedError bool
	}{
		{
			name:          "Valid URL HTTPS .com",
			body:          `{"original_url": "https://example.com"}`,
			expectedURL:   "https://example.com",
			expectedError: false,
		},
		{
			name:          "Valid URL HTTPS .net",
			body:          `{"original_url": "https://example.net/"}`,
			expectedURL:   "https://example.net/",
			expectedError: false,
		},
		{
			name:          "Invalid URL",
			body:          `{"original_url": "htp://invalid-url"}`,
			expectedURL:   "",
			expectedError: true,
		},
		{
			name:          "Unreachable URL",
			body:          `{"original_url": "https://dnaqd2ownasdawk.com"}`,
			expectedURL:   "",
			expectedError: true,
		},
		{
			name:          "Random String",
			body:          `{"original_url": "dnalwisgkawd"}`,
			expectedURL:   "",
			expectedError: true,
		},
		{
			name:          "Empty URL field",
			body:          `{}`,
			expectedURL:   "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/", nil)
			req.Body = io.NopCloser(strings.NewReader(tt.body))

			u, err := decodeAndVerifyOriginalURL(req)

			if (err != nil) != tt.expectedError {
				t.Errorf("decodeAndVerifyOriginalURL() error = %v, expectedError %v", err, tt.expectedError)
				return
			}

			if err == nil && u.String() != tt.expectedURL {
				t.Errorf("decodeAndVerifyOriginalURL() = %v, expected %v", u.String(), tt.expectedURL)
			}
		})
	}
}
