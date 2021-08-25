package main

func main() {
	// 在slice或map中存储指针或者包含指针的值。
	var x = 10
	var ls []*int
	ls = append(ls, &x)
	var y int
	var mp map[string]*int
	mp["y"] = &y
}
