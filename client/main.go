package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {

	port := flag.Int("p", 24680, "ListenUDP port")
	flag.Parse()

	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: *port,
	})
	if err != nil {
		panic(err)
	}

	go func() {
		body := make([]byte, 1024)
		for {
			n, addr, err := listen.ReadFromUDP(body)
			if err != nil {
				panic((err))
			}
			fmt.Printf("cli recv: %s, addr: %s \n", string(body[:n]), addr.String())
		}
	}()

	for {
		_, err := listen.WriteTo([]byte("ping"), &net.UDPAddr{
			IP:   net.IPv4(97, 64, 17, 204), // net.IPv4(97, 64, 17, 204), // net.IPv4(192, 168, 8, 8), // net.IPv4(132,232,112,242),
			Port: 13579,
		})
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		time.Sleep((time.Second * 30))
	}
}
