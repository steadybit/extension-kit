#!/usr/bin/env bash
# SPDX-License-Identifier: MIT
# SPDX-FileCopyrightText: 2026 Steadybit GmbH
#
# Regenerates the Go lexer/parser/visitor in internal/gen/ from the grammar in grammar/.
#
# The grammar is a *copy* of the platform's source of truth at
#   platform/query-language/src/main/antlr4/com/steadybit/ql/
# A drift check in the platform repository fails if the two diverge. Never edit the copy here to
# change the language — change it there and re-copy, or the two parsers will disagree.
#
# Usage: ./generate.sh
set -euo pipefail

# Keep in sync with <antlr.version> in the platform's pom.xml. The generated code and the
# github.com/antlr4-go/antlr/v4 runtime in go.mod must come from the same 4.13.x line.
ANTLR_VERSION="4.13.2"
# SHA-256 of antlr-4.13.2-complete.jar as published on antlr.org. Update it together with
# ANTLR_VERSION. This jar is executed, so it is pinned rather than merely downloaded: without the
# check, anything that can answer for that URL — or a cache poisoned by an earlier run — decides
# what code generates our parser.
ANTLR_SHA256="eae2dfa119a64327444672aff63e9ec35a20180dc5b8090b7a6ab85125df4d76"

cd "$(dirname "$0")"

CACHE_DIR="${ANTLR_CACHE_DIR:-${TMPDIR:-/tmp}/steadybit-antlr}"
JAR="${CACHE_DIR}/antlr-${ANTLR_VERSION}-complete.jar"

# shasum on macOS, sha256sum on most Linux distributions.
sha256_of() {
  if command -v sha256sum >/dev/null 2>&1; then
    sha256sum "$1" | cut -d' ' -f1
  else
    shasum -a 256 "$1" | cut -d' ' -f1
  fi
}

if [[ ! -f "${JAR}" ]]; then
  echo "Downloading ANTLR ${ANTLR_VERSION} to ${JAR}"
  mkdir -p "${CACHE_DIR}"
  # --proto '=https' and --proto-redir '=https' keep the transfer on HTTPS even if the server
  # answers with a redirect; -L on its own would happily follow one to plain HTTP.
  curl -sSfL --proto '=https' --proto-redir '=https' \
    -o "${JAR}" "https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar"
fi

ACTUAL_SHA256="$(sha256_of "${JAR}")"
if [[ "${ACTUAL_SHA256}" != "${ANTLR_SHA256}" ]]; then
  # Remove it, so a corrupted download does not poison every later run from the cache.
  rm -f "${JAR}"
  echo "ANTLR ${ANTLR_VERSION} checksum mismatch — refusing to run it." >&2
  echo "  expected ${ANTLR_SHA256}" >&2
  echo "  actual   ${ACTUAL_SHA256}" >&2
  exit 1
fi

OUT="${PWD}/internal/gen"
rm -rf "${OUT}"
mkdir -p "${OUT}"

# ANTLR mirrors the input file's relative path under -o, so invoke it from inside grammar/ with an
# absolute output directory to get flat output. The lexer must be generated first: the parser
# grammar declares `tokenVocab=QueryLanguageLexer` and needs the resulting .tokens file, which -lib
# points at.
(
  cd grammar
  for g in QueryLanguageLexer.g4 QueryLanguageParser.g4; do
    java -jar "${JAR}" \
      -Dlanguage=Go \
      -package gen \
      -visitor \
      -no-listener \
      -lib "${OUT}" \
      -o "${OUT}" \
      "${g}"
  done
)

# The .interp/.tokens files are build intermediates, not Go sources.
rm -f internal/gen/*.interp internal/gen/*.tokens

gofmt -w internal/gen

echo "Generated:"
ls -1 internal/gen
