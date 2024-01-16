package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 13579,
	})
	if err != nil {
		panic(err)
	}
	body := make([]byte, 1024)
	for {
		n, addr, err := listen.ReadFromUDP(body)
		if err != nil {
			panic((err))
		}

		fmt.Printf("svr recv: %s, addr: %s \n", string(body[:n]), addr.String())

		_, err = listen.WriteTo([]byte("echo"), addr)
		if err != nil {
			panic((err))
		}
	}
}
