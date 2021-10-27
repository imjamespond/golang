# [Quich Start](https://grpc.io/docs/languages/go/quickstart/)
```
# 安装protoc命令
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
export PATH="$PATH:$(go env GOPATH)/bin" # protoc 加入path
```
[Download examples](https://github.com/grpc/grpc-go/archive/v1.35.0.zip)

```
ProtoPath=./test/hello/proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ${ProtoPath}/hello.proto
```
``paths=source_relative`` 设置执行目录与source相对路径