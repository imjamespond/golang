package main

import (
	"4d-qrcode/controller"
	"log"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Thanks to otraore for the code example
// https://gist.github.com/otraore/4b3120aa70e1c1aa33ba78e886bb54f3

func main() {

	g := engine()
	g.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func engine() *gin.Engine {
	g := gin.New()

	g.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))
	g.POST("/gen-qrcode", controller.GenQrcode)
	g.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	return g
}
