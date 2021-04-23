package main

import (
	"log"
	"test-gin-auth/controller"
	"test-gin-auth/service"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Thanks to otraore for the code example
// https://gist.github.com/otraore/4b3120aa70e1c1aa33ba78e886bb54f3

func main() {

	service.ConnectToMysql()

	g := engine()
	g.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func engine() *gin.Engine {
	g := gin.New()
	g.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	g.POST("/login", controller.Login)
	g.GET("/logout", controller.Logout)
	g.GET("/version", Version)

	private := g.Group("/private")
	private.Use(controller.AuthRequired)
	{
		private.GET("/me", controller.Me)
		private.GET("/status", controller.Status)
	}
	return g
}
