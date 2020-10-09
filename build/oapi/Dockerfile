FROM golang:1.15-alpine3.12 as base
RUN apk update && apk upgrade && \ 
    apk add --no-cache bash git openssh protobuf
ENV GOBIN /go/bin
RUN export PATH=$PATH:$GOBIN
RUN go get -u -v \
    github.com/deepmap/oapi-codegen/cmd/oapi-codegen

FROM base as builder
ENV GOBIN /go/bin
RUN export PATH=$PATH:$GOBIN; \
    mkdir -p /tmp/gen/
ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" skipcache
ADD ./api/dist/ ./api/
RUN oapi-codegen --generate types,chi-server,spec ./api/openapi.yaml  > /tmp/gen/openapi.gen.go
