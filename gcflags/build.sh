Main=${1-main.go}
# 其中 -m 表示进行内存分配分析，-l 表示避免程序内联
go build -gcflags "-m=2 -l" ${Main}