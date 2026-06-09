# Contributing

## Prerequisites

To work on this library locally, you need:

- [Go](https://go.dev/dl/) 1.25 or later
- [GNU Make](https://www.gnu.org/software/make/)

## Tasks

The `Makefile` in the project root contains commands to easily run common admin tasks.

| Command        | Meaning                                                                     |
|----------------|-----------------------------------------------------------------------------|
| `$ make tidy`  | Format all code using `go fmt` and tidy the `go.mod` file.                  |
| `$ make audit` | Run `go vet`, `staticcheck`, execute all tests and verify required modules. |

## Releasing

To make a new release, do the following:

 1. Update the `CHANGELOG.md`
 2. Commit and push the changelog changes.
 3. Set the tag `git tag -a vX.X.X -m vX.X.X`
 4. Push the tag.

## Generating New Test Certificates

```
openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out cert.pem \
  -keyout key.pem \
  -addext "subjectAltName = DNS:localhost" \
  -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"
```

## Contributor License Agreement (CLA)

In order to accept your pull request, we need you to submit a CLA. You only need to do this once. If you are submitting a pull request for the first time, just submit a Pull Request and our CLA Bot will give you instructions on how to sign the CLA before merging your Pull Request.

All contributors must sign an [Individual Contributor License Agreement](https://github.com/steadybit/.github/blob/main/.github/cla/individual-cla.md).

If contributing on behalf of your company, your company must sign a [Corporate Contributor License Agreement](https://github.com/steadybit/.github/blob/main/.github/cla/corporate-cla.md). If so, please contact us via office@steadybit.com.

If for any reason, your first contribution is in a PR created by other contributor, please just add a comment to the PR
with the following text to agree our CLA: "I have read the CLA Document and I hereby sign the CLA".
