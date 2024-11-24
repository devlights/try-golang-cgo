package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const int BUF_SIZE = 11;

void setBuf(char *buf) {
	strncpy(buf, "helloworld", BUF_SIZE-1);
}
*/
import "C"
import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

func main() {
	var (
		buf      = make([]byte, int(C.BUF_SIZE))
		bufPtr   = unsafe.Pointer(&buf[0]) //明示的に先頭要素のポインタを渡す
		charPtr  = (*C.char)(bufPtr)       // (void *) --> (char *)
		printBuf = func() {
			fmt.Println(string(buf))
			fmt.Println(hex.Dump(buf))
		}
	)

	printBuf()
	{
		// C側の宣言では (char *) を引数に要求しているため、unsafe.Pointer を *C.char にキャストして渡す.
		// unsafe.Pointer は、 (void *) を表すので、任意のポインタ型にキャスト可能.
		C.setBuf(charPtr)
	}
	printBuf()
}
