package main

import "C"

//当使用export的时候，在同一个文件中就不能再定义其它的c函数了，不然会报错。
//使用export导出函数给c语言调用。

//export GoAdd
func GoAdd(a, b int) int {
	return a + b
}
