package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *cStr = "hello C  World";

void pStr(const char *m) {
	printf("%s (char *)\n", m);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	//
	// Goの文字列をC側の文字列にするには C.CString()  を使う
	// C側の文字列をGoの文字列にするには C.GoString() を使う
	//

	//
	// Go -> C
	//
	var (
		goStr = "hello Go World"
		cStr  = C.CString(goStr)
	)
	defer C.free(unsafe.Pointer(cStr))

	C.pStr(cStr)

	//
	// C -> Go
	//
	var (
		cStr2  = C.cStr
		goStr2 = C.GoString(cStr2)
	)

	fmt.Printf("%[1]v (%[1]T)\n", goStr2)
}
