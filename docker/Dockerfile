FROM golang:1.14.2
LABEL maintainer="huynhquangthao@gmail.com"

RUN go get -u golang.org/x/lint/golint

ENV APP_HOME  /opencensus-prometheus
WORKDIR $APP_HOME
COPY go.mod go.sum $APP_HOME/
RUN GO111MODULE=on go mod download

COPY . $APP_HOME/
