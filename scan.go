package scan

import (
	"arping"
	"encoding/binary"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Isport(port string) bool {
	var validator = regexp.MustCompile("^((6553[0-5])|(655[0-2][0-9])|(65[0-4][0-9]{2})|(6[0-4][0-9]{3})|([1-5][0-9]{4})|([0-5]{0,5})|([0-9]{1,4}))$")
	return validator.MatchString(port)
}

func domainchecker(hostname string) bool {
	domain := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
 ]{2,3})$`)
	return domain.MatchString(hostname)
}

func Get_ip(ip string) []string {
	var (
		err    error
		ips    []net.IP
		all_ip []string
	)
	if domainchecker(ip) {
		ips, err = net.LookupIP(ip)
		for _, ip := range ips {
			if ipv4 := ip.To4(); ipv4 != nil {
				all_ip = append(all_ip, ipv4.String())
			}

		}
		if err != nil {
			ip = ""
		}
	} else if net.ParseIP(ip) == nil {
		ip = ""
	} else {
		all_ip = append(all_ip, ip)
	}
	return all_ip
}

func socket(ip string, port int) (socket string) {
	socket = ip + ":" + strconv.Itoa(port)
	return socket
}

func Tcp_scan(ip string, port int,timeout time.Duration) int {
	connection, err := net.DialTimeout("tcp", socket(ip, port), timeout)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			Tcp_scan(ip, port, timeout)
		} else {
			return 0
		}
	}

	defer connection.Close()
	return port

}
func Cdirgetter(cidr string) ([]string, error) {
	var hosts []string
	_, subnet, err := net.ParseCIDR(cidr)
    if err != nil {
        print("Please Input a valid CIDR in this format (192.168.1.1/24, 10.0.0.0/8)")
        os.Exit(0)
    }
	mascara := binary.BigEndian.Uint32(subnet.Mask)
	fAddr := binary.BigEndian.Uint32(subnet.IP)
	lAddr := (fAddr & mascara) | (mascara ^ 0xffffffff)
	for i := fAddr; i <= lAddr; i++ {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		ips := ip.String()
		hosts = append(hosts, ips)
	}
	return hosts, err
}
func Arpscan_lan(ips string) (string, string) {
	ip := net.ParseIP(ips)
	arping.SetTimeout(500 * time.Millisecond)
	HwAddr, _, err := arping.Ping(ip)
	mac := HwAddr.String()
	if err == arping.ErrTimeout {
		return mac, ""
	} else if err != nil {
		if strings.Contains(err.Error(), "operation not") {
			print("Please run as root\n")
			os.Exit(1)
		} else if strings.Contains(err.Error(), "ip+net") {
			return mac, "Fail in net resources occurred Running again" + "\n"
			Arpscan_lan(ips)

		} else if strings.Contains(err.Error(), "no usable interface found") {
			print("You put CIDR of another net OR Try Run same as root\n")
			os.Exit(0)
			return mac, "Probably you put a CIDR outside ur net" + "\n"

		} else {

			print("Running again: Unknown Error succedeed try run program in root\n")
			os.Exit(1)
			return mac, "Running again: Unknown Error succedeed try run program in root\n"

		}
	} else {
		return mac, ips
	}
	return mac, "Error"
}
