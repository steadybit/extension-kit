/*
 * Copyright 2026 steadybit GmbH. All rights reserved.
 */

package exthttp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBumpRevisionChangesRevision(t *testing.T) {
	before := Revision()
	BumpRevision()
	after := Revision()

	assert.NotEqual(t, before, after, "revision must change after BumpRevision")
}

func TestRevisionStableWithoutBump(t *testing.T) {
	assert.Equal(t, Revision(), Revision(), "revision must be stable without a bump")
}

func TestRegisterIndexHandlerEmitsEtagAndServes304(t *testing.T) {
	handler := IfNoneMatchHandler(Revision, GetterAsHandler(func() map[string]string {
		return map[string]string{"hello": "world"}
	}))

	rec := httptest.NewRecorder()
	handler(rec, httptest.NewRequest(http.MethodGet, "/", nil), nil)

	require.Equal(t, http.StatusOK, rec.Code)
	etag := rec.Header().Get("ETag")
	require.NotEmpty(t, etag, "ETag header must be set on a 200 response")
	assert.Equal(t, Revision(), etag)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("If-None-Match", etag)
	rec = httptest.NewRecorder()
	handler(rec, req, nil)

	assert.Equal(t, http.StatusNotModified, rec.Code, "matching If-None-Match must return 304")
}
