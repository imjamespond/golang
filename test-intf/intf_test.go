package testintf

import (
	"fmt"
	"testing"
)

type IFoobar interface {
	Hello(val string) (rt string)
}

type FoobarImpl struct {
}

func (f FoobarImpl) Hello(val string) (rt string) {
	return "Hello," + val
}

// var (
// 	foobar IFoobar = FoobarImpl{}
// )

func sayHello() IFoobar {
	return &FoobarImpl{}
}

func Test1(t *testing.T) {
	fmt.Println(sayHello().Hello("Jacob"))
}
