package main

import (
	"flag"
	"fmt"
	"http-server/controllers"
	ffmpeg_service "http-server/service/ffmpeg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	host := flag.String("h", "127.0.0.1", "host to listen")
	port := flag.String("p", "3000", "port to serve on")
	flag.Parse()

	ffmpegService := ffmpeg_service.GetInstance()
	ffmpegService.Start()

	r := gin.New()
	r.StaticFS("/statics", http.Dir("./"))
	api := r.Group("/api")
	api.POST("/download-m3u8", controllers.DownloadM3u8)
	api.POST("/convert-video", controllers.ConvertVideo)
	api.POST("/kill", controllers.Kill)
	api.GET("/log", controllers.GetLog)
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", *host, *port),
		Handler:        r,
		MaxHeaderBytes: 1 << 22,
	}
	err := s.ListenAndServe()
	log.Fatal(err, nil)
}
