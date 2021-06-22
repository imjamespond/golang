# protoc
```
# go get -u github.com/golang/protobuf
go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2
# 1.3.2以上生成go代码版本不兼容？
```
`这一步会自动在gopath/bin目录下面生成一个protobuf-gen-go,`  
[解压protoc-3.17.3-osx-x86_64.zip](https://github.com/protocolbuffers/protobuf/releases)`到gopath/bin, 执行`  
`protoc -I=./model --go_out=plugins=grpc:. book.proto` 

# svr
```
go test ./test/svr_test.go
debug TestCli in vscode 
```