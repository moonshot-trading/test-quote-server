FROM golang:1.9.3-alpine3.7
MAINTAINER tzinck

run apk update && apk upgrade && apk add --no-cache bash git openssh

RUN mkdir -p /go/src/github.com/moonshot-trading/test-quote-server
ADD . /go/src/github.com/moonshot-trading/test-quote-server
RUN go get github.com/moonshot-trading/test-quote-server
RUN go install github.com/moonshot-trading/test-quote-server

ENTRYPOINT /go/bin/test-quote-server