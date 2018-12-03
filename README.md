# Boy Oh Boy Stock Quote Fetcher

A simple microservice conforming to the [Open Microservice Specification](https://microservice.guide/) (OMS) that fetches the most recent trade price (in dollars) for a provided ticker symbol e.g. GOOG, MSFT, PVTL

## Building and Running

There are two ways to build this, directly as a binary via:

```
go build
./boyohboy fetch "{ \"ticker\": \"PVTL\" }"
```

or via Docker:

```
docker build -t boyohboy:latest .
docker run boyohboy /boyohboy fetch "{ \"ticker\": \"PVTL\" }"
```

## Running the tests

The integration test suite is written using [Ginkgo](https://onsi.github.io/ginkgo/). If you have ginkgo installed, they can be run via:

```
ginkgo -r
```

Alternatively, with fewer features, using `go test`:

```
go test ./...
```
