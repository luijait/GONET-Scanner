package ports
func Ports() map[int]string {
    //Based in well known ports
    ports := map[int]string {
        1:"echo",
        9:"WOL",
        20:"ftp data",
        21:"ftp control",
        22:"ssh",
        23:"telnet",
        25:"smtp",
        43:"whois",
        49:"TACACS",
        53:"DNS",
        67:"BOOTP",
        69:"TFTP",
        70:"Gopher",
        71:"NETRJS",
        80:"http",
        81:"TorPark",
        82:"TorPark",
        88:"Kerberos",
        123:"NTP",
        135:"RPC",
        443:"https",
        445:"Microsoft-ds, Samba",
        465:"SMTP over TLS",
        514:"Syslog",
        520:"RIP",
        521:"RIPng",
        540:"UUCP",
        543:"klogin",
        544:"kshell",



    }
    return ports
}
