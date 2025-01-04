package main

/*
#include <stdlib.h>

extern void my_printf(const char *);
extern void func1();
*/
import "C"
import "unsafe"

func init() {
	println(">> gofuncs.init() called")
	m["gofuncs"] = true
}

func func2() {
	msg := C.CString("helloworld")
	defer C.free(unsafe.Pointer(msg))

	C.my_printf(msg)
	C.func1()
}
