package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const int SKIP_BYTES = 6;

// cgo側でポインタ演算後のアドレスで処理
void p(const void *v) {
	char *m = (char *)v;
	printf("%s\n", m);
}

// 自分でポインタ演算して処理
void p2(const void *v, const int skip) {
	char *m = (char *)v;
	printf("%s\n", m+skip);
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
		offsetPtr = unsafe.Pointer(uintptr(cStrPtr) + uintptr(C.SKIP_BYTES))
	)

	C.p(offsetPtr)
	C.p2(cStrPtr, C.SKIP_BYTES)
}
