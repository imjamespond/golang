# go build or run main in subfolder 报错
```
# command-line-arguments
jwt-server/server.go:25:27: undefined: NewDummyTokenStore
```
` go run jwt-server/server.go jwt-server/tokenStore.go `  
` go build -o jwt-server/server jwt-server/*.go `

