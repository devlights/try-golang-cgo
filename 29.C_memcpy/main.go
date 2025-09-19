package main

/*
#cgo CFLAGS: -Wall -std=c99
#cgo LDFLAGS:

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"strings"
	"unsafe"
)

//export mymemcpy
func mymemcpy(dst *C.char, src *C.char, n C.size_t) {
	//
	// memcpy関数のシグネチャは
	//   void *memcpy(void *dest, const void *src, size_t n);
	// となっている。
	//
	// cgo側から見ると以下のように見える。
	//   func memcpy(dst unsafe.Pointer, src unsafe.Pointer, n C.size_t) (r1 unsafe.Pointer)
	//
	// (void *) は、cgoでは、unsafe.Pointer となるため
	// 指定する値の型が (char *) つまり、(*C.char) の場合は
	// unsafe.Pointer(*C.char)に変換して渡す。
	//

	// 元文字列を大文字に変換してmemcpyする
	//   - C.GoString()とC.CString()については、サンプル「06.C_GoString」を参照。

	var (
		goStr = strings.ToUpper(C.GoString(src))
		cStr  = unsafe.Pointer(C.CString(goStr))
	)
	defer C.free(cStr)

	var (
		pDst = unsafe.Pointer(dst)
	)
	C.memcpy(pDst, cStr, n)
}

func main() {}
