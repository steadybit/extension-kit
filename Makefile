# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	gofmt -l .
# -unreachable=false: ANTLR's Go target ends every generated rule function with a
# `goto errorExit` after the `return`, to keep the label used in the rules that need it
# (extquery/internal/gen, 46 occurrences). That code is generated and checked in, so it cannot
# be edited, and the trick cannot be stripped without breaking the rules that do use the label.
# Every other vet analyzer still runs, on every package.
	go vet -unreachable=false ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-SA1019,-ST1000,-U1000,-ST1003 ./...
	TZ=UTC go test -race -vet=off -coverprofile=coverage.out ./...
	go mod verify
