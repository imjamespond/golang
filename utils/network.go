package utils

import (
	"codechiev/utils"
	"net"
)

func IpList() []*net.IP {
	ifaces, err := net.Interfaces()
	utils.FatalIf(err)

	ips := make([]*net.IP, 0)
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
			// fmt.Println(ip)
			ips = append(ips, &ip)
		}
	}
	return ips
}

// Get preferred outbound ip of this machine
func Ip(conn net.Conn) net.IP {
	// conn, err := net.Dial("udp", "8.8.8.8:80")
	// utils.FatalIf(err)
	// defer conn.Close()

	localAddr := conn.LocalAddr().(*net.TCPAddr) //.(*net.UDPAddr)

	return localAddr.IP
}
