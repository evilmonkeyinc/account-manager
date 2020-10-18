#!/bin/sh
cd api && npm run bundle && cd ..
rm -rf gen/go
mkdir -p gen/go
docker run --name openapi-generator --rm -v "${PWD}:/local" swaggerapi/swagger-codegen-cli-v3 generate \
    -i /local/api/dist/openapi.yaml \
    -l go-server \
    -o /local/gen/go