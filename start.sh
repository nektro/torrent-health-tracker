#!/usr/bin/env bash

set -e
go test
go build
./torrent-health-tracker \
