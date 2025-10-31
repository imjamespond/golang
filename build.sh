rsync -av --delete ~/workspace/webapp/_test/test-webrtc/dist/webrtc/ ./webrtc/
GOOS=windows GOARCH=386 go build -o webrtc-server.exe
