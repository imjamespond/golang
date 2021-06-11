package foo

import (
	"fmt"

	"example.com/test/foobar/bar"
	"example.com/test/log"
)

func PrintFoo() {
	log.DummyLog()
	bar.PrintBar()
	fmt.Println("test foo...")
}
