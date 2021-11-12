# [Setup walk](https://github.com/lxn/walk) 
```
go get github.com/akavel/rsrc
rsrc -manifest test.manifest -o rsrc.syso

# 去掉cmd
go build -ldflags="-H windowsgui"
```