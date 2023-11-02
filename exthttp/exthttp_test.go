/*
 * Copyright 2023 steadybit GmbH. All rights reserved.
 */

package exthttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequestTimeoutHeaderAware(t *testing.T) {
	tests := []struct {
		name                 string
		requestTimeoutHeader string
		wantedStatusCode     int
		wantsDeadline        bool
	}{
		{
			name:                 "Should apply no timeout if no header is set",
			requestTimeoutHeader: "",
			wantedStatusCode:     200,
			wantsDeadline:        false,
		},
		{
			name:                 "Should apply no timeout if invalid header is set",
			requestTimeoutHeader: "foobar",
			wantedStatusCode:     200,
			wantsDeadline:        false,
		},
		{
			name:                 "Should apply timeout if valid header is set",
			requestTimeoutHeader: "0.5",
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
			if len(tt.requestTimeoutHeader) > 0 {
				req.Header.Set("Request-Timeout", tt.requestTimeoutHeader)
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

func TestIfNoneMatchHandler(t *testing.T) {
	tests := []struct {
		name                     string
		IfNoneMatchRequestHeader string
		ETag                     string
		wantedEtagResponseHeader string
		wantedStatusCode         int
	}{
		{
			name:                     "should return 200 and no header",
			IfNoneMatchRequestHeader: "",
			ETag:                     "",
			wantedEtagResponseHeader: "",
			wantedStatusCode:         200,
		},
		{
			name:                     "should return 200 and etag header",
			IfNoneMatchRequestHeader: "",
			ETag:                     "abcdef",
			wantedEtagResponseHeader: "abcdef",
			wantedStatusCode:         200,
		},
		{
			name:                     "should return 304 and etag header",
			IfNoneMatchRequestHeader: "abcdef",
			ETag:                     "abcdef",
			wantedEtagResponseHeader: "",
			wantedStatusCode:         304,
		},
		{
			name:                     "should return 200 and etag header",
			IfNoneMatchRequestHeader: "abcdef",
			ETag:                     "ghijkl",
			wantedEtagResponseHeader: "ghijkl",
			wantedStatusCode:         200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			if len(tt.IfNoneMatchRequestHeader) > 0 {
				req.Header.Set("If-None-Match", tt.IfNoneMatchRequestHeader)
			}

			rr := httptest.NewRecorder()
			IfNoneMatchHandler(func() string {
				return tt.ETag
			}, func(w http.ResponseWriter, r *http.Request, body []byte) {
				w.WriteHeader(200)
			})(rr, req, nil)

			if status := rr.Code; status != tt.wantedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantedStatusCode)
			}

			if etag := rr.Header().Get("ETag"); etag != tt.wantedEtagResponseHeader {
				t.Errorf("handler returned wrong etag header: got %v want %v", etag, tt.wantedEtagResponseHeader)
			}
		})
	}
}
