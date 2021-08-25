package main

func main() {
	// 向chan中发送数据的指针或者包含指针的值。
	// 因为编译器此时不知道值什么时候会被接收，因此只能放入堆中
	x := 1
	c := make(chan int, 1)
	c <- x // x不发生逃逸，因为只是复制的值
	c1 := make(chan *int, 1)
	y := 2
	c1 <- &y // y逃逸，因为地址传入了chan
	z := 3
	pz := &z // z不逃逸，因为是确定性析构
	*pz += 1
}
