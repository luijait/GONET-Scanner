#!/bin/bash

GOROOT=$(go env | grep "GOROOT" | sed 's/\GOROOT="//g' | sed 's/\"//g')
cp scan.go $GOROOT/scan/
cp ports.go $GOROOT/ports/
