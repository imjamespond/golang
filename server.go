package main

import (
	"4d-qrcode/controller"
	"4d-qrcode/util"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var startOK func()

func main() {
	fmt.Println(syscall.Getpid())

	// startOK = service_pdf.RunPdfkit

	// engine := engine()
	// engine.Use(gin.Logger())
	// if err := engine().Run(":8080"); err != nil {
	// 	log.Fatal("Unable to start:", err)
	// }

	NotifyWithoutContext()
}

func engine() *gin.Engine {
	eng := gin.New()

	eng.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))
	eng.POST("/gen-qrcode", controller.GenQrcode)
	eng.NoRoute(func(ctx *gin.Context) {
		ctx.File("./frontend/dist/index.html")
	})

	return eng
}

func launch() *http.Server {
	router := engine() // gin.Default()
	router.GET("/test-timeout", func(c *gin.Context) {
		time.Sleep(30 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	router.GET("/test-interrupt", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT) // or stopChan <- syscall.SIGINT
		c.String(http.StatusOK, "SIGINT has been sent...")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		// if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// 	log.Fatalf("listen: %s\n", err)
		// }
		l, err := net.Listen("tcp", "localhost:18080")
		util.FatalIf(err)
		if startOK != nil {
			startOK()
		}
		err = srv.Serve(l)
		util.FatalIf(err)
	}()

	return srv
}

func NotifyWithContext() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := launch()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func NotifyWithoutContext() {
	srv := launch()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
