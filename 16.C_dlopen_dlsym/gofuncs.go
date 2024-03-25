package main

/*
extern size_t call_my_strlen(void *fn, const char *s);
*/
import "C"
import (
	"log"
	"unsafe"

	"github.com/devlights/try-golang-cgo/16.C_dlopen_dlsym/dlopen"
)

//export my_strlen
func my_strlen(cStr *C.char) C.size_t {
	var (
		handle unsafe.Pointer
		symbol unsafe.Pointer
		result C.size_t
	)

	handle = dlopen.OpenLib("clib/libclib.so")
	defer dlopen.CloseLib(handle)

	symbol = dlopen.GetSym(handle, "my_strlen")
	result = C.call_my_strlen(symbol, cStr)

	log.Printf("[Go][インターセプト][my_strlen] result=%v", result)

	return result
}
