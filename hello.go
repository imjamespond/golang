package main

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1 -I .
#cgo LDFLAGS: -L${SRCDIR} -lhello
#include <stdio.h>
#include "hello.h"

#if defined(CGO_OS_WINDOWS)
    const char* os = "windows";
#else
    const char* os = "not win32";
#endif
void hello(){
    printf("hello: %s\n", os);
}
void printint(int v) {
    printf("printint: %d\n", v);
}
*/
import "C"

func hello() {
	v := 42
	C.printint(C.int(v))
	C.hello()
	C.hello_world()
}
