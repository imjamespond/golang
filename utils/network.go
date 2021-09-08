package utils

import (
	"codechiev/utils"
	"fmt"
	"net"
)

func IpList() {
	ifaces, err := net.Interfaces()
	utils.FatalIf(err)
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		utils.FatalIf(err)
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			fmt.Println(ip)
		}
	}
}

// Get preferred outbound ip of this machine
func Ip(conn net.Conn) net.IP {
	// conn, err := net.Dial("udp", "8.8.8.8:80")
	// utils.FatalIf(err)
	// defer conn.Close()

	localAddr := conn.LocalAddr().(*net.TCPAddr) //.(*net.UDPAddr)

	return localAddr.IP
}
