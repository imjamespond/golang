package main

func foo() []int {
	return []int{1, 2, 3}
}

func main() {
	ls := foo() // 发生逃逸
	// ls := []int{1, 2, 3}
	ls = append(ls, 1)
}
