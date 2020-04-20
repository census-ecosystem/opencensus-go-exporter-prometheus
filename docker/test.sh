#!/usr/bin/env bash

golint  \
&& go test ./... -race -p 1  \
&& go clean -testcache ./...
