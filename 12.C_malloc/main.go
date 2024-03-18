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
		bufPtr unsafe.Pointer
		szBuf  = C.size_t(11)
	)

	// malloc()を呼び出しヒープメモリを確保してもらう.
	// 内部では malloc() をラップしたヘルパー関数が呼び出される.
	bufPtr = C.malloc(szBuf)
	defer C.free(bufPtr)

	C.memset(bufPtr, C.int(0), szBuf)

	//
	// 文字列をコピー
	//
	var (
		cStr    = C.CString("helloworld")
		cStrPtr = unsafe.Pointer(cStr)
	)
	defer C.free(cStrPtr)

	C.memcpy(bufPtr, cStrPtr, szBuf)

	//
	// C側でちゃんと見れるか確認
	//
	C.p(bufPtr)
}
