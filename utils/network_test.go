package utils

import (
	"codechiev/utils"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"testing"
)

func TestIp(t *testing.T) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	utils.FatalIf(err)
	defer conn.Close()
	ip := Ip(conn)
	log.Println(ip)
}

func TestIpList(t *testing.T) {
	ips := IpList()
	for _, ip := range ips {
		data, _ := json.Marshal(ip)
		fmt.Println(string(data))
	}
}
