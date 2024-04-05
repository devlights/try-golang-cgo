package main

/*
#include <stdlib.h>

int cfunc1(char *name) {
	if (name == NULL) {
		return -1;
	}

	return 0;
}
*/
import "C"
import (
	"log"
	"unsafe"
)

func init() {
	log.SetFlags(0)
}

func main() {
	var (
		cNULL  = unsafe.Pointer(nil)
		cName  = (*C.char)(cNULL)
		cName2 = C.CString("helloworld")
	)
	defer C.free(unsafe.Pointer(cName2))

	log.Printf("NULL     ==> %d", int(C.cfunc1(cName)))
	log.Printf("NOT NULL ==> %d", int(C.cfunc1(cName2)))
}
