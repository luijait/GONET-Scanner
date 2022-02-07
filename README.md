# GO/NET Scanner
<div id="image" align="center">
  <img src="https://user-images.githubusercontent.com/60628803/152819380-79b7f954-faca-4363-a44b-86a5eb3a68cb.png" >
  </div>
  
 ---
 
 <div id="badges" align="center">
  <img src="https://img.shields.io/badge/%40author-luijait.es-informational">
  <img src="https://img.shields.io/github/repo-size/luijait/GONET-Scanner?label=Size">
  <img src="https://img.shields.io/github/languages/top/luijait/GONET-Scanner?label=go">
</div>


---

# ScreenShots

<img src="https://user-images.githubusercontent.com/60628803/152821931-a2678f6c-c383-4939-9040-938d5f01defd.png">
<img src="https://user-images.githubusercontent.com/60628803/152824097-301c66b1-5248-4995-b1ee-1f509c1cb184.png">

---

# Install
```
chmod +x install.sh
./install.sh [as root]
```

---


# Usage
```
[ARGUMENTS]

-ar CIDR: ARP Discovery
-ar CIDR -s: Scan ports in all hosts discovered
-ap: Scan to 65535 Ports
-pr MINPORT MAXPORT: Define Port Range to Scan
-1000: Scan Top 1000 ports (like nmap)
-t: Set Timeout (in milliseconds)

[EXAMPLES]

go run scannerport.go -ap <IP>: Allports TCP Scan
go run scannerport.go <IP> Default Scan 0-1024 ports
go run scannerport.go  -ar 192.168.0.1/24 <IP>: ARP Ping Scan ALL local Subnet
go run scannerport.go <IP> -pr <MINPORT> <MAXPORT>
go run scannerport.go -ar 192.168.1.1/24 -s
go run scannerport.go -1000 192.168.1.1
go run scannerport.go -t 100 192.168.1.1
Example: go run scannerport.go -ar 192.168.1.1/24 (will send an arp ping to every host of net to discover if is it up)
Example: go run scannerport.go google.com -1000 (Will resolve google.com + Will scan top 1000 ports)
Example: go run scannerport.go 192.168.0.1 -pr 100 3000 (will scan every port in these range you must put first minor port)
```
