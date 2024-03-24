package main

/*
#include <stdlib.h>

extern size_t call_clib_my_strlen(void *fn, const char *s);
*/
import "C"
import (
	"log"
	"unsafe"

	"github.com/devlights/try-golang-cgo/16.C_dlopen_dlsym/dlopen"
)

//export clib_my_strlen
func clib_my_strlen(cStr *C.char) C.size_t {
	var (
		handle unsafe.Pointer
		symbol unsafe.Pointer
		result C.size_t
	)

	println("1")
	handle = dlopen.OpenLib("clib/libclib.so")
	defer dlopen.CloseLib(handle)

	println("2")
	println(handle == nil)
	symbol = dlopen.GetSym(handle, "clib_my_strlen")
	result = C.call_clib_my_strlen(symbol, cStr)

	log.Printf("[Go][strlen] result=%v", result)

	return result
}

func main() {}
