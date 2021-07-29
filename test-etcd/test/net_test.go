package test

import (
	"codechiev/utils"
	"fmt"
	"net"
	"testing"
)

func TestNetIp(t *testing.T) {
	ifaces, err := net.Interfaces()
	utils.PanicIf(err)
	// handle err
	for _, i := range ifaces {
		fmt.Println("ifaces: ", i.Name, i.HardwareAddr.String())
		addrs, err := i.Addrs()
		utils.PanicIf(err)
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				fmt.Printf("*net.IPNet: ")
				ip = v.IP
			case *net.IPAddr:
				fmt.Printf("*net.IPAddr: ")
				ip = v.IP
			}
			// process IP address
			fmt.Println(ip)
		}
	}
}

// Get preferred outbound ip of this machine
func TestNetGetOutboundIP(t *testing.T) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	utils.FatalIf(err)
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	fmt.Println(localAddr.IP)
}
