#!/bin/sh
docker rm $(docker ps -a -q)
docker image prune --force