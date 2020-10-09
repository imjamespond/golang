package main

import (
	"log"
  "net/http"
  "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	
	"example.com/greetings"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	msg, err := greetings.Hello("Foobar")
	if err != nil {
		log.Fatal(err)
	}
  return c.String(http.StatusOK, msg)
}