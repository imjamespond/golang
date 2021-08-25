package main

import (
	"os"
	"strconv"
)

func main() {
	ls := []int{1, 2, 3}
	ls = append(ls, 4) // 确定性的，不逃逸，编译期间可以知道
	var n int
	n, _ = strconv.Atoi(os.Args[1]) // 输入数据后，则结果不可知，因此可能逃逸
	ls1 := []int{1, 2, 3}
	for i := 0; i < n; i++ {
		ls1 = append(ls1, 1)
	}
}
