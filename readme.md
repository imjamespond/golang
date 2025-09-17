```bash
# init 
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
go get github.com/oapi-codegen/runtime/types
go get github.com/gin-gonic/gin

# gen code
oapi-codegen -package main -generate types,gin-server -o generated.go openapi.yaml
# run main
go run .
# go live swagger-ui.html
```