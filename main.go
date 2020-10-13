package main

import (
	"log"

	"github.com/webview/webview"
)

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetSize(800, 600, webview.HintNone)
	w.SetTitle("Hello")
	w.Bind("hello", func() {
		log.Println("hello")
	})
	w.Bind("noop", func() string {
		log.Println("hello")
		return "hello"
	})
	w.Bind("add", func(a, b int) int {
		return a + b
	})
	w.Bind("quit", func() {
		//w.Terminate()
	})
	w.Navigate(`data:text/html,
			<!doctype html>
			<html>
				<body><button onclick="hello()">hello?</button></body>
				<script>
					window.onload = function() {
						//document.body.innerText = ` + "`hello, ${navigator.userAgent}`" + `;
						noop().then(function(res) {
							console.log('noop res', res);
							add(1, 2).then(function(res) {
								console.log('add res', res);
								quit();
							});
						});
					};
				</script>
			</html>
		`)
	w.Run()
}
