/*
 * Copyright 2023 steadybit GmbH. All rights reserved.
 */

package exthttp

import (
	"github.com/steadybit/extension-kit/extutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequestTimeoutHeaderAware(t *testing.T) {
	tests := []struct {
		name                 string
		requestTimeoutHeader *string
		wantedStatusCode     int
		wantsDeadline        bool
	}{
		{
			name:                 "Should apply no timeout if no header is set",
			requestTimeoutHeader: nil,
			wantedStatusCode:     200,
			wantsDeadline:        false,
		},
		{
			name:                 "Should apply no timeout if invalid header is set",
			requestTimeoutHeader: extutil.Ptr("foobar"),
			wantedStatusCode:     200,
			wantsDeadline:        false,
		},
		{
			name:                 "Should apply timeout if valid header is set",
			requestTimeoutHeader: extutil.Ptr("0.5"),
			wantedStatusCode:     503,
			wantsDeadline:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			if tt.requestTimeoutHeader != nil {
				req.Header.Set("Request-Timeout", *tt.requestTimeoutHeader)
			}

			rr := httptest.NewRecorder()
			RequestTimeoutHeaderAware(handler(t, tt.wantsDeadline)).ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantedStatusCode)
			}
		})
	}
}

func handler(t *testing.T, wantsDeadline bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		w.WriteHeader(200)
		if _, ok := r.Context().Deadline(); wantsDeadline != ok {
			t.Errorf("Expected request context to have a deadline, but it didn't")
		}
	}
}
