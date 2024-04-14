package main

/*
#cgo CFLAGS: -g -O0 -Wall -Wextra

#include <stdio.h>
#include <stdlib.h>

void myPrint(const char *msg) {
	printf("%s\n", msg);
}
*/
import "C"
import "unsafe"

func main() {
	var (
		goStr = "helloworld"
		cStr  = C.CString(goStr)
	)
	defer C.free(unsafe.Pointer(cStr))

	C.myPrint(cStr)
}
