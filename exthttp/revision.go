/*
 * Copyright 2026 steadybit GmbH. All rights reserved.
 */

package exthttp

import (
	"strconv"
	"sync/atomic"
	"time"
)

// The revision is the ETag for the combined extension index endpoint ("/"). It changes on process
// start (seeded with a nanosecond startup nonce) and whenever a kit SDK registers or clears a
// describable element (via BumpRevision). This lets the agent cache the index response and skip the
// per-element describe calls on a matching If-None-Match, without relying on a process restart to
// invalidate the cache.
//
// The value is kept as a plain, unquoted string so it round-trips through the agent's
// If-None-Match/ETag handling byte-for-byte, matching the historical startedAt behavior.
var (
	revisionSeed    = strconv.FormatInt(time.Now().UnixNano(), 36)
	revisionCounter atomic.Uint64
)

// Revision returns the current index revision string. It changes on process start and whenever any
// kit registers/unregisters a describable element, so it is a safe ETag for the combined index.
func Revision() string {
	return revisionSeed + "-" + strconv.FormatUint(revisionCounter.Load(), 10)
}

// BumpRevision advances the index revision. Kit SDKs call this from their Register*/Clear* functions
// whenever the set of registered describable elements changes.
func BumpRevision() {
	revisionCounter.Add(1)
}

// RegisterRevisionedHandler registers a handler that serves the getter's result at path, tagging the
// response with the current Revision() as its ETag so it supports conditional GET (a matching
// If-None-Match yields 304). Use this instead of hand-rolling IfNoneMatchHandler with a
// process-start timestamp, e.g. for the combined extension index.
func RegisterRevisionedHandler[T any](path string, getter func() T) {
	RegisterHttpHandler(path, IfNoneMatchHandler(Revision, GetterAsHandler(getter)))
}
