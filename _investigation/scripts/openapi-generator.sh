#!/bin/sh
cd api && npm run bundle && cd ..
rm -rf gen
docker run --name openapi-generator --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/api/dist/openapi.yaml \
    -g go-server \
    --enable-post-process-file \
    -o /local/gen/go