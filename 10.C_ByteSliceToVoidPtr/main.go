package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const int BUF_SIZE = 11;

void setBuf(void *buf) {
	strncpy((char *)buf, "helloworld", BUF_SIZE);
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
		printBuf = func() {
			fmt.Println(string(buf))
			fmt.Println(hex.Dump(buf))
		}
	)

	printBuf()

	// C側の宣言では (void *) を引数に要求しているため、unsafe.Pointer を渡す.
	// unsafe.Pointer は、 (void *) を表す.
	C.setBuf(bufPtr)

	printBuf()
}
