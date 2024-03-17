package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void pStr(const char *m) {
	printf("%s\n", m);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	//
	// cgoにてC言語側の文字列を扱いたい場合は
	// C.CStringを利用する。
	//
	// C.CStringは、内部でmallocを使ってヒープメモリを
	// 確保しているため、利用後は必ず C.free() にてメモリを開放する必要がある。
	//
	// C.CStringは、C言語の世界の (char *) を表す。
	// (unsafe.Pointer() は、C言語の世界の (void *) を表す)
	//
	var (
		goStr = "this is go string"
		cStr  = C.CString(goStr)
	)

	// 必ず C.free() で開放する
	defer C.free(unsafe.Pointer(cStr))

	fmt.Printf("%[1]v (%[1]T)\n", goStr)
	fmt.Printf("%[1]v (%[1]T)\n", cStr)
	C.pStr(cStr)
}
