#!/bin/bash
if [ "$(id -u)" != "0" ]; then
		echo
		echo "This script must be run as root." 1>&2
		echo
		exit 1
fi

if [-z "$GOROOT"]; then
    echo "GOROOT env variable not set. Please set it before executing this script."
fi

if [!-d "$GOROOT/src/arping"]; then
    git clone https://github.com/j-keck/arping -o $GOROOT/src/arping
fi

if [!-d "$GOROOT/src/scan"];then
    mkdir $GOROOT/src/scan
fi

if [!-d "$GOROOT/src/ports"]; then
    mkdir $GOROOT/src/ports
fi

cp scan.go $GOROOT/src/scan
cp ports.go $GOROOT/src/ports

go build scannerPort.go