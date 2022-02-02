#!/bin/bash
if [-z "$GOROOT"]; then
    echo "GOROOT env variable not set. Please set it before executing this script."
fi

mkdir $GOROOT/scan
mkdir $GOROOT/ports

cp scan.go $GOROOT/scan
cp ports.go $GOROOT/ports

go build scannerPort.go