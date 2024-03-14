package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sample.h"

typedef struct {
	int value1;
	char value2[20];
} ST1;

void myprintf(char *s) {
	printf("%s\n", s);
}

void myprintf2(char *s) {
	prt(s);
}

size_t mysizeof(ST1 *v) {
	return sizeof(*v);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	s := C.CString("helloworld")
	defer C.free(unsafe.Pointer(s))

	//
	// 可変長引数を持つ関数は呼べない。
	//
	// printf()は、可変長引数を持っているため、そのまま呼べない。
	// 以下はコンパイルエラーとなる。
	// (cgo: 02.Restrictions/main.go:21:2: unexpected type: ...)
	//
	//C.printf(unsafe.Pointer(s))

	// 代わりに cgo ヘッダ部に定義した関数を呼び出すようにする
	C.myprintf(s)

	//
	// マクロは呼べない
	//
	// 以下の C.prt は、C側に定義されているマクロを呼び出そうとしているが
	// コンパイルエラーとなる。
	// (./main.go:35:2: could not determine kind of name for C.prt)
	//
	// 同様に cgo 側にラッパー関数を用意して呼び出すのは問題無い。
	//
	//C.prt(s)
	C.myprintf2(s)

	//
	// sizeof演算子は使えない
	//
	// sizeof演算子はコンパイル時にデータ型のサイズを決定するC言語の演算子なので
	// cgoからは利用できない。
	//
	// 同様に cgo 側にラッパー関数を用意して呼び出すのは問題無い。
	//
	var st1 C.ST1

	st1.value1 = 999

	buf := make([]byte, 20)
	copy(buf, []byte("helloworld"))

	cBuf := C.CBytes(buf)
	defer C.free(cBuf)
	C.memcpy(unsafe.Pointer(&st1.value2), cBuf, C.ulong(len(buf)))

	szSt1 := C.mysizeof((*C.ST1)(unsafe.Pointer(&st1)))

	fmt.Printf("%+v (%d bytes)\n", st1, szSt1)
}
