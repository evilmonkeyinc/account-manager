#!/bin/sh
docker build -t generator:0.0.1 -f ./build/Dockerfile .
docker run --name gen-api-proto generator:0.0.1 /bin/bash
rm -rf ./gen/
docker cp gen-api-proto:/tmp/gen/. gen/
docker rm gen-api-proto