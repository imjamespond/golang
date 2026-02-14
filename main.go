package main

import (
	"flag"
	"fmt"
	"jamespond/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	host := flag.String("h", "127.0.0.1", "host to listen")
	port := flag.String("p", "3000", "port to serve on")
	flag.Parse()

	handler := gin.New()
	handler.StaticFS("/statics", http.Dir("./"))
	api := handler.Group("/api")
	api.POST("/action", controllers.Action)
	handler.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", *host, *port),
		Handler:        handler,
		MaxHeaderBytes: 1 << 22,
	}
	err := server.ListenAndServe()
	log.Fatal(err, nil)
}
