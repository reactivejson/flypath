# flypath# Flight Path API


![Maintainer](https://img.shields.io/badge/maintainer-MohamedAly-blue)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/reactivejson/flypath.svg)](https://github.com/reactivejson/flypath)
[![Go Reference](https://pkg.go.dev/badge/github.com/reactivejson/flypath)](https://pkg.go.dev/badge/github.com/reactivejson/flypath)
[![Go](https://github.com/reactivejson/flypath/actions/workflows/go.yml/badge.svg)](https://github.com/reactivejson/flypath/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/reactivejson/flypath)](https://goreportcard.com/report/github.com/reactivejson/flypath)
![](https://img.shields.io/github/license/reactivejson/flypath.svg)

Go microservice API that can help understand and track how a particular person's flight path may be queried.
The API accepts a request that includes a list of flights, which are defined by a source and destination airport code.
These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

## Endpoints
The API has a single endpoint:

POST /calculate
This endpoint accepts a JSON payload containing a list of flights and returns the complete flight path as a JSON response.

Request Payload
The request payload should be a JSON object contains array of flight objects. Each flight object should have two fields, source and destination, which represent the source and destination airports.

Example request payload:
````json
[
  {
    "source": "IND",
    "destination": "EWR"
  },
  {
    "source": "SFO",
    "destination": "ATL"
  },
  {
    "source": "GSO",
    "destination": "IND"
  },
  {
    "source": "ATL",
    "destination": "GSO"
  }
]

````

Response Payload
The response payload is a JSON object with a single field flight_path that is an object of airport codes representing the complete flight path, starting with the starting airport and ending with the ending airport.

Example response payload:

````json
{"flight_path":
  { "source":"SFO",
    "destination":"EWR"
  }
}
````

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
    ├── helm
    │   └── <helm chart files>
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

### Helm & K8S
Helm charts to deploy this micro-service in a Kubernetes platform
We generate the container image and reference it in a Helm chart

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

### Deploy with Helm chart in a Kubernetes environment
```bash
 helm upgrade --namespace neo --install vwap-engine chart/vwap-engine -f your-custom-values.yml
```

## Test coverage

Test coverage is checked as a part of test execution with the gotestsum tool.

Test coverage is checked for unit tests and integration tests.

Coverage report files are available and stored as `*coverage.txt` and are also imported in the SonarQube for easier browsing.


## golangci-lint

In the effort of reducing errors and improving the overall quality of code, golangci-lint is run as a part of the pipeline. Linting is run for the services and packages that have changes since the previous green build (in master) or previous commit (in local or review).

Any issues found by golangci-lint for the changed code will lead to a failed build.

golangci-lint rules are configured in `.golangci.yml`.

### Logic
This Go code defines a package called pathFinder, which contains a type called Flight that represents a flight from a source airport to a destination airport. The Flight type has two fields, Source and Destination, which are both strings.

The code also defines a type called tracker, which has three fields: a slice of Flight objects, and two slices of strings called starts and ends. The starts slice contains the source airports of all the flights, while the ends slice contains the destination airports of all the flights.

The New function is a constructor for the tracker type. It takes a slice of Flight objects as an argument and returns a pointer to a new tracker object. The New function initializes the starts and ends slices with the same length as the fl slice passed in, and sets them to the corresponding source and destination airports from the fl slice.

The PathFinder method on the tracker type returns a Flight object that represents the complete flight path from the starting airport to the ending airport. It does this by calling the setSlices method to update the starts and ends slices, and then calling the DiffSlice function to find the airports that appear in one slice but not the other. The Source field of the returned Flight object is set to the airport that appears in the starts slice but not the ends slice, and the Destination field is set to the airport that appears in the ends slice but not the starts slice.

The setSlices method sets the starts and ends slices based on the Source and Destination fields of the Flight objects in the fl slice.

The DiffSlice function takes two slices of strings as arguments and returns the first string in the first slice that does not appear in the second slice. It does this by creating a map of the strings in the second slice, and then iterating through the strings in the first slice to find the first string that does not appear in the map.
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
