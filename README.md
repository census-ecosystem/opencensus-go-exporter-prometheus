> **Warning**
>
> OpenCensus and OpenTracing have merged to form [OpenTelemetry](https://opentelemetry.io), which serves as the next major version of OpenCensus and OpenTracing.
>
> OpenTelemetry has now reached feature parity with OpenCensus, with tracing and metrics SDKs available in .NET, Golang, Java, NodeJS, and Python. **All OpenCensus Github repositories, except [census-instrumentation/opencensus-python](https://github.com/census-instrumentation/opencensus-python), will be archived on July 31st, 2023**. We encourage users to migrate to OpenTelemetry by this date.
>
> To help you gradually migrate your instrumentation to OpenTelemetry, bridges are available in Java, Go, Python, and JS. [**Read the full blog post to learn more**](https://opentelemetry.io/blog/2023/sunsetting-opencensus/).

# OpenCensus Go Prometheus Exporter

[![Build Status](https://travis-ci.org/census-ecosystem/opencensus-go-exporter-prometheus.svg?branch=master)](https://travis-ci.org/census-ecosystem/opencensus-go-exporter-prometheus) [![GoDoc][godoc-image]][godoc-url]

Provides OpenCensus metrics export support for Prometheus.

## Installation

```
$ go get -u contrib.go.opencensus.io/exporter/prometheus
```

[godoc-image]: https://godoc.org/contrib.go.opencensus.io/exporter/prometheus?status.svg
[godoc-url]: https://godoc.org/contrib.go.opencensus.io/exporter/prometheus
