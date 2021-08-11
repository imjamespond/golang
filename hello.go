package main

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1 -I .
#cgo LDFLAGS: -L${SRCDIR} -lhello
#include <stdio.h>
#include <stdlib.h> //cstring 用到?
#include "hello.h" //注意只能是调c的.h,不能为.hpp

#if defined(CGO_OS_WINDOWS)
    const char* os = "windows";
#else
    const char* os = "not win32";
#endif

int GoAdd(int a, int b);

const char* hello(char *str){
    printf("hello %s on %s\n", str, os);
    return "bar";
}
void printint(int v) {
    printf("printint: %d\n", GoAdd(v,v));
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func hello() {
	v := 42
	C.printint(C.int(v))
	foo := C.CString("foo")
	bar := C.hello(foo)
	fmt.Println("hello from", C.GoString(bar))
	C.hello_world()

	defer C.free(unsafe.Pointer(foo))
}
