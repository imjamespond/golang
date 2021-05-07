package main

import (
	"fmt"
	"net"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func startServer() {

	g := gin.Default()

	g.Use(gin.Logger())

	g.Use(static.Serve("/app", static.LocalFile("../web/dist", false)))

	if err := g.Run(":8080"); err != nil {
		panic(err)
	}
}

func StartWeb() {

	go startServer()

	fmt.Println("\nTrying to connect to web server:")
	for {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", "8080"), time.Second*5)
		if err != nil {
			fmt.Println("Connecting error:", err)
		}
		if conn != nil {
			fmt.Println("Connected!")
			// defer conn.Close()
			conn.Close()
			break
		} else {
			time.Sleep(time.Second >> 1)
		}
	}

}
