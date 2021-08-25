package main

import "time"

func main() {
	x := 10
	go func(x *int) {
		*x += 1
	}(&x) // 捕获的瞬间，x没有移动到heap上，但是整个闭包移动到了heap上，因此x也跟随闭包被移动到heap上了
	time.Sleep(time.Second * 2)
}
