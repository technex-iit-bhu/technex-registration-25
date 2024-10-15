#!/bin/bash

find . -name "*.go" -type f -exec gofmt -w {} \;

echo "All Go files have been formatted."
