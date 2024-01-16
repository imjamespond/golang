export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
export
go build -ldflags="-s -w" -o svr ./server/main.go 