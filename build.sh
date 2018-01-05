#!/bin/sh

export GITHUB_TOKEN="$GORELEASE_GITHUB_TOKEN"
goreleaser --rm-dist