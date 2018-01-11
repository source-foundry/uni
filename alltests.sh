#!/usr/bin/env bash

go build .
./uni --version
pytest
go test -v ./...

