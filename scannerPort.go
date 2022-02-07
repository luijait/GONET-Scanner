package main

import (
	"fmt"
	"os"
	"ports"
	"scan"
	"strconv"
	"time"
	"strings"
)

const (
	MAXPORT int = 65535
	MINPORT int = 0
	BPORT   int = 1024
)

var (
	ip           string
	ips          []string
	parameter    string
	cidr         string
	arguments    []string
	parameters   []string
	hosts_online []string
	notmatch     int
	finish_port  int
	timeouts	 int
	start_port   int
	err          error
	scansubnet   bool
	mk_ports     bool
	
	ports_list   map[int]string
    mac string
)

func getService(ports map[int]string, port int) string {
	service := ports[port]
	return service
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func port_parser(start_port int, finish_port int) (int, int) {
	if start_port > finish_port {
		aux := finish_port
		finish_port = start_port
		start_port = aux
	}

	return start_port, finish_port
}
func remove(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func scan_type(args []string) (int, int) {
	start_port := MINPORT
	finish_port := BPORT
	for i := range args {
		if args[i] == "-ap" {
			finish_port = MAXPORT
		} else if args[i] == "-pr" {
			if scan.Isport(args[i+1]) && scan.Isport(args[i+2]) {
				start_port, err = strconv.Atoi(args[i+1])
				finish_port, err = strconv.Atoi(args[i+2])
				if err != nil {
					print(err)
					man_menu()
				} else {
					start_port, finish_port = port_parser(start_port, finish_port)
				}
			}
		}
	}
	return start_port, finish_port
}
func isnotempty(s string) bool {
	return len(s) > 0
}

func man_menu() {
    fmt.Print("[ARGUMENTS]\n\n")
    fmt.Print("-ar CIDR: ARP Discovery\n-ar CIDR -s: Scan ports in all hosts discovered\n-ap: Scan to 65535 Ports\n-pr MINPORT MAXPORT: Define Port Range to Scan\n-1000: Scan Top 1000 ports (like nmap)\n-t: Set Timeout (in milliseconds)\n\n")
    fmt.Print("[EXAMPLES]\n\n")
	fmt.Print("go run scannerport.go -ap <IP>: Allports TCP Scan\n")
	fmt.Print("go run scannerport.go <IP> Default Scan 0-1024 ports\n")
	fmt.Print("go run scannerport.go  -ar 192.168.0.1/24 <IP>: ARP Ping Scan ALL local Subnet\n")
	fmt.Print("go run scannerport.go <IP> -pr <MINPORT> <MAXPORT>\n")
	fmt.Print("go run scannerport.go -ar 192.168.1.1/24 -s\n")
	fmt.Print("go run scannerport.go -1000 192.168.1.1")
	fmt.Print("go run scannerport.go -t 100 192.168.1.1")
	fmt.Print("Example:", " go run scannerport.go -ar 192.168.1.1/24 (will send an arp ping to every host of net to discover if is it up)\n")
	fmt.Print("Example:", " go run scannerport.go google.com -1000 (Will resolve google.com + Will scan top 1000 ports)")
	fmt.Print("Example: go run scannerport.go 192.168.0.1 -pr 100 3000 (will scan every port in these range you must put first minor port)")
	os.Exit(0)

}

func Args(arguments []string) (timeout int, mkports bool, scansubnet bool, cidr string, args []string, ip string) {
	timeout = 100 //Default Timeout
	var parameters = []string{
		"-ar",
		"-ap",
		"-pr",
		"-s",
		"-1000",
		"-t",
	}
	remove(arguments, 0)
	for i := range arguments {
		notmatch = 0
		if len(scan.Get_ip(arguments[i])) > 0 {
			ips = scan.Get_ip(arguments[i])
			remove(arguments, i)
		} else {
			for j := range parameters {
				if arguments[i] != parameters[j] {
					notmatch++
				} else if arguments[i] == "-ar" {
					parameters = append(parameters, arguments[i+1])
					cidr = arguments[i+1]
					if arguments[i+2] == "-s" {
						scansubnet = true
					} 
					remove(arguments, i+1)
				} else if arguments[i] == "-1000" {
					mk_ports = true
					
				}
				if isNumeric(arguments[i]) {
					notmatch = notmatch - 1
				}
			}
			if notmatch == len(parameters) {
				man_menu()
			}
			
		}
		if arguments[i] == "-t" {
			timeout, err = strconv.Atoi(arguments[i+1]) 
			if err != nil {
				man_menu()
			}
		}
	}
	if len(ips) == 0 && !isnotempty(cidr) {
		man_menu()
	}
	return timeout ,mk_ports, scansubnet, cidr, arguments, ip
}
func printer(mac string, host string, port int, service string) {
	if !isnotempty(service) {
		service = "Not Found"
	}
	if port == 0 {
		print("")
	} else if port != MAXPORT+1 {

		print(fmt.Sprintf("%d\tOpen\t"+service+"\n", port))
	} else {
		if strings.Contains(host, "Running") {
			print(host + "\n")
			fmt.Print("|HOST SCANNED|\t|STATE|\t|Mac Addresses|\n")
		} else {
			print(fmt.Sprintf("%s\tOnline\t"+mac+"\n", host))
		}

	}
}

func main() {

	ports_list = ports.Ports()
	timeouts ,mk_ports, scansubnet, cidr, arguments, ip = Args(os.Args)
	timeout := time.Duration(timeouts)
	if isnotempty(cidr) {
		hosts, err := scan.Cdirgetter(cidr)
		if err != nil {
			man_menu()
		}
		fmt.Print("|HOST SCANNED|\t|STATE|\t|Mac Addresses|\n")
		for i := range hosts {
			mac, host := scan.Arpscan_lan(hosts[i])
			if isnotempty(host) {
				printer(mac, host, MAXPORT+1, "")
				hosts_online = append(hosts_online, host)
			}
		}
		if scansubnet {
			start_port, finish_port = scan_type(arguments)
			for i := range hosts_online {
				print("Currently scanning host: " + hosts_online[i] + "\n")

				fmt.Printf("PORT\tSTATUS\tService\n")
				for j := start_port; j < finish_port; j++ {
						printer("", hosts_online[i], scan.Tcp_scan(hosts_online[i], j, timeout * time.Millisecond), getService(ports_list, j))
				}
			}
		}

	} else if mk_ports {
		for i := range ips {
			ports_1000 := ports.Top1000ports() 
			ip = ips[i]
			fmt.Printf("Host to Scan: " + ip + "\n")
			fmt.Printf("PORT" + "\t" + "STATUS " + "\t" + "Service" + "\n")
			for j := range ports_1000 {
				
					printer("", ip, scan.Tcp_scan(ip, ports_1000[j], timeout * time.Millisecond), getService(ports_list, ports_1000[j]))

}
} 
} else {
		start_port, finish_port = scan_type(arguments)
		for i := range ips {
			ip = ips[i]
			fmt.Printf("Host to Scan: " + ip + "\n")
			fmt.Printf("PORT" + "\t" + "STATUS " + "\t" + "Service" + "\n")
			for j := start_port; j < finish_port; j++ {
				printer("", ip, scan.Tcp_scan(ip, j, timeout * time.Millisecond), getService(ports_list, j))	
			}
		}
	} 
}
