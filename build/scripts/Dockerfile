FROM golang:1.15-alpine3.12 as base
RUN apk update && apk upgrade && \ 
    apk add --no-cache bash git openssh protobuf
ENV GOBIN /go/bin
RUN export PATH=$PATH:$GOBIN
RUN go get -u -v \
github.com/golang/protobuf/proto \
github.com/golang/protobuf/protoc-gen-go \
google.golang.org/grpc \
github.com/googleapis/gnostic \
github.com/googleapis/gnostic-grpc \
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN git clone https://github.com/protocolbuffers/protobuf.git

FROM base as builder
ENV GOBIN /go/bin
RUN export PATH=$PATH:$GOBIN; \
mkdir -p /tmp/gen/google; \
cp -r ./src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/. /tmp/gen/google/; \
cp -r ./protobuf/src/google/. /tmp/gen/google/
ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" skipcache
ADD ./api/ ./api/
RUN gnostic --grpc-out=/tmp/gen/ ./api/openapi.yaml; \ 
protoc --proto_path=/tmp/gen/ --go_out=plugins=grpc:/tmp/gen/ --grpc-gateway_out=logtostderr=true,paths=source_relative:/tmp/gen/ /tmp/gen/openapi.proto
RUN rm -rf /tmp/gen/google/
