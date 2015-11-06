#!/usr/bin/env bash

export GOOS=windows
export GOARCH=386

go build -v -o bin/registry.exe