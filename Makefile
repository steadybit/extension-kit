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
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-SA1019,-ST1000,-U1000,-ST1003 ./...
	TZ=UTC go test -race -vet=off -coverprofile=coverage.out ./...
	go mod verify
