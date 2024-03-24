//go:build run

package main

/*
#cgo LDFLAGS: -ldl

#include <dlfcn.h>
#include <stdlib.h>

typedef size_t (*strlen_t)(const char *);

strlen_t getStrlenFunc();

size_t callstrlen(const char *s);
*/
import "C"

//export strlen
func strlen(cstr *C.char) C.size_t {
	return C.callstrlen(cstr)
}

func main() {}
