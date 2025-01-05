package main

/*
#cgo CFLAGS: -I../lib
#cgo LDFLAGS: -L../lib -lmylib

#include <stdlib.h>
#include "lib.h"
*/
import "C"
import "unsafe"

func main() {
	var (
		goStr = "hello world"
		cStr  = C.CString(goStr)
	)
	defer C.free(unsafe.Pointer(cStr))

	C.myPrint(cStr)
}
