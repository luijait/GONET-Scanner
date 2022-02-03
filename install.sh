#!/bin/bash
if [ "$(id -u)" != "0" ]; then
		echo
		echo "This script must be run as root." 1>&2
		echo
		exit 1
fi
apt install golang
Dir=$(go env | grep "GOROOT" | sed 's/\GOROOT="//g' | sed 's/\"//g')
[ ! -d "$Dir/src/arping" ] && git clone https://github.com/j-keck/arping -o $Dir/src/arping

[ ! -d "$Dir/src/scan" ] && mkdir "$Dir/src/scan" || echo "Already exists scan"


[ ! -d "$Dir/src/ports" ] && mkdir "$Dir/src/ports" || echo "Already exists ports"

cp scan.go "$Dir/src/scan" 
cp ports.go "$Dir/src/ports"

go build scannerPort.go
