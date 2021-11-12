package main

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1 -I .
#cgo LDFLAGS: -L${SRCDIR} -lhello
#include <stdio.h>
#include <stdlib.h> //cstring 用到?'
#include <string.h>
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
const void* abc(void *str){
	char dest[50];
	memset(dest, 0, sizeof dest);
	memcpy(dest, str, 3);
	printf("say %s\n", (const char*)dest);
	return "Efg";
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
	abc := C.CBytes([]byte{'a', 'b', 'c'})
	bar := C.hello(foo)
	efg := C.abc(abc)
	fmt.Println("from hello:", C.GoString(bar))
	fmt.Println("from abc:", string(C.GoBytes(efg, 3)))
	C.hello_world()

	defer C.free(unsafe.Pointer(foo))
	defer C.free(abc)
}
