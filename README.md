# OpenCensus Go Prometheus Exporter

[![Build Status](https://travis-ci.org/census-ecosystem/opencensus-go-exporter-prometheus.svg?branch=master)](https://travis-ci.org/census-ecosystem/opencensus-go-exporter-prometheus) [![GoDoc][godoc-image]][godoc-url]

Provides OpenCensus metrics export support for Prometheus.

## Installation

```
$ go get -u github.com/hqt/opencensus-pushgateway-exporter
```

## Running
```go
exporter, err := NewPushGatewayExporter(serviceName, getGatewayURL(), 300*time.Millisecond)
if err != nil {
	panic(err)
}
exporter.Run()
defer exporter.Close()
```

## Testing:
- Run the pushgateway container:
```bash
docker run -d -p 9091:9091 prom/pushgateway
```

- Run tests using the local environment (start the pushgateway container first):
```bash
make test
```

- Run tests using the Docker environment:
```bash
make ci-test
```

[godoc-image]: https://godoc.org/contrib.go.opencensus.io/exporter/prometheus?status.svg
[godoc-url]: https://godoc.org/contrib.go.opencensus.io/exporter/prometheus
