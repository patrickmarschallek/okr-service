#!/bin/sh
# The following command will produce a statically linked Go binary without debugging (dwarf) information.
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

# build docker container.
docker build -t pmarschallek/okr-service:latest .

# start docker container
docker run -d -P pmarschallek/okr-service