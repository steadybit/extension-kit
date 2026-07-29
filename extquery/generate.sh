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

cd "$(dirname "$0")"

CACHE_DIR="${ANTLR_CACHE_DIR:-${TMPDIR:-/tmp}/steadybit-antlr}"
JAR="${CACHE_DIR}/antlr-${ANTLR_VERSION}-complete.jar"

if [[ ! -f "${JAR}" ]]; then
  echo "Downloading ANTLR ${ANTLR_VERSION} to ${JAR}"
  mkdir -p "${CACHE_DIR}"
  curl -sSfL -o "${JAR}" "https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar"
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
