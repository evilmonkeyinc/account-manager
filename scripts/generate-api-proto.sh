#!/bin/sh
cd api && npm run bundle && cd ..
docker build -t protobuf-generator -f ./build/proto/Dockerfile .
docker run --name gen-api-proto protobuf-generator /bin/bash
rm -rf gen/
mkdir -p gen/openapi
docker cp gen-api-proto:/tmp/gen/. gen/openapi/
docker rm gen-api-proto