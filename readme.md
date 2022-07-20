# Upgrade golang

## vscode, 删除原有的gopath内容
```
# vscode 代码提示
go install -v golang.org/x/tools/gopls@latest
# vscode 调试
go install -v github.com/go-delve/delve/cmd/dlv@latest 
```

## Issues
```
compile: version "go1.16.4" does not match go tool version "go1.18.4"
```
查看export,GOROOT指向旧版,`更新后要完全重启vscode`