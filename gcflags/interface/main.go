package main

type foo interface {
	fooFunc()
}

type foo1 struct {
	foo
}

func (f1 foo1) fooFunc() {}

type foo2 struct{}

func (f2 *foo2) fooFunc() {}

func main() {
	// interface类型的GC，涉及使用interface类型转换并调用对应的方法的时候，都会发生内存逃逸，给出代码示例：
	var f foo
	f = foo1{}
	f.fooFunc() // 调用方法时，发生逃逸，因为方法是动态分配的
	f = &foo2{}
	f.fooFunc()
}
