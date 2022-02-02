if [ "$(id -u)" != "0" ]; then
		echo
		echo "This script must be run as root." 1>&2
		echo
		exit 1
	fi

echo "For Working This script you need GOROOT enviroment variable set"
mkdir $GOROOT/src/scan 
cp scan.go $GOROOT/src/scan/scan.go
git clone https://github.com/j-keck/arping -o $GOROOT/src/arping
