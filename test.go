package main

import (
	"unsafe"
)

/*
#include <stdlib.h>

typedef int(*test_func)(int a, int b);
*/
import "C"

//export add
func add(a, b C.int) C.int {
	aGo := int(a)
	bGo := int(b)
	res := aGo + bGo
	return C.int(res)
}

//export hello
func hello(name *C.char) *C.char {
	data := C.GoString(name)
	cs := C.CString("Hello " + data + "!")
	return cs
}

//export freePoint
func freePoint(cs *C.void) {
	C.free(unsafe.Pointer(cs))
}

//export test
func test(cb C.test_func) int {
	// rs := cb(1, 2)
	// fmt.Println(rs)
	return 3
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
