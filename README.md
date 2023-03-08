# flypath# Flight Path API


![Maintainer](https://img.shields.io/badge/maintainer-MohamedAly-blue)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/reactivejson/flypath.svg)](https://github.com/reactivejson/flypath)
[![Go Reference](https://pkg.go.dev/badge/github.com/reactivejson/flypath)](https://pkg.go.dev/badge/github.com/reactivejson/flypath)
[![Go](https://github.com/reactivejson/flypath/actions/workflows/go.yml/badge.svg)](https://github.com/reactivejson/flypath/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/reactivejson/flypath)](https://goreportcard.com/report/github.com/reactivejson/flypath)
![](https://img.shields.io/github/license/reactivejson/flypath.svg)

Go microservice API that can help us understand and track how a particular person's flight path may be queried.
The API accepts a request that includes a list of flights, which are defined by a source and destination airport code.
These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

### Project layout

This layout is following pattern:

```text
flypath
└───
    ├── .github
    │   └── workflows
    │     └── go.yml
    ├── cmd
    │   └── main.go
    ├── internal
    │   └── pathFinder.go
    ├── build
    │   └── Dockerfile
    ├── Makefile
    ├── README.md
    └── <source packages>
```

## Setup

### Getting started
flypath is available in github
[Flypath](https://github.com/reactivejson/flypath)

```shell
go get github.com/reactivejson/flypath
```

#### Run
```shell
go run cmd/main.go
```

#### Build
```shell
make build
```
#### Testing
```shell
make test
```
### Build docker image:

```bash
make docker-build
```
This will build this application docker image so-called flypath

## Test coverage

Test coverage is checked as a part of test execution with the gotestsum tool.

Test coverage is checked for unit tests and integration tests.

Coverage report files are available and stored as `*coverage.txt` and are also imported in the SonarQube for easier browsing.


## golangci-lint

In the effort of reducing errors and improving the overall quality of code, golangci-lint is run as a part of the pipeline. Linting is run for the services and packages that have changes since the previous green build (in master) or previous commit (in local or review).

Any issues found by golangci-lint for the changed code will lead to a failed build.

golangci-lint rules are configured in `.golangci.yml`.

### Requirements

- Go 1.18 or newer [https://golang.org/doc/install](https://golang.org/doc/install)
- Docker 18.09.6 or newer

### Variable names
Commonly used one letter variable names:

- i for index
- r for reader
- w for writer
- c for client
- flypath Flight Path


## License

Apache 2.0, see [LICENSE](LICENSE).
