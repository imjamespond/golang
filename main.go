package main

import (
	"log"
	"net/http"
	"embed"
)

/* This is a Go compiler directive. It tells the compiler to embed all files and directories inside the webrtc folder into the variable content. */
//go:embed webrtc/*
var content embed.FS

func main() {
	// 指定 WebRTC 文件的目录
	// 请将 "path/to/your/webrtc/dir" 替换为您的实际目录路径
	webrtcDir := "./webrtc"

	// 创建一个文件服务器，用于服务静态文件
	// StripPrefix 用于移除 URL 前缀，以便文件服务器能正确找到文件
	// http.Handle("/webrtc/", http.StripPrefix("/webrtc/", http.FileServer(http.Dir(webrtcDir))))
	http.Handle("/webrtc/", http.FileServer(http.FS(content)))

	// 启动 HTTP 服务器
	log.Println("Serving WebRTC files from:", webrtcDir)
	log.Println("Server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}