package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void p(const void *v) {
	char *m = (char *)v;
	printf("%s\n", m);
}
*/
import "C"
import "unsafe"

func main() {
	var (
		goStr   = "hello Go World"
		cStr    = C.CString(goStr)
		cStrPtr = unsafe.Pointer(cStr)
	)
	defer C.free(cStrPtr)

	C.p(cStrPtr)

	// uintptr に変換することでポインタ演算が可能となる
	// 演算後を再度 unsafe.Pointer にする
	//
	// 以下はメモリアドレスを6バイト進めたポインタを取得している
	var (
		offsetPtr = unsafe.Pointer(uintptr(cStrPtr) + 6)
	)

	C.p(offsetPtr)
}
