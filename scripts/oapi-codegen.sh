#!/bin/sh
cd api && npm run bundle && cd ..
docker build -t oapi-codegen -f ./build/oapi/Dockerfile .
docker run --name oapi-codegen oapi-codegen /bin/bash
rm -rf gen/
docker cp oapi-codegen:/tmp/gen/. gen/
docker rm oapi-codegen