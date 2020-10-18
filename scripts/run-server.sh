#!/bin/sh
docker stop account-manager
docker rm account-manager
docker build -t account-manager -f ./build/package/server/Dockerfile .
docker run -p 8080:8081 --name account-manager -d account-manager