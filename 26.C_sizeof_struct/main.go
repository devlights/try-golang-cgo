package main

/*
#include <stdio.h>
#include <stdlib.h>

struct A {
	int value;
};

void A_setvalue(struct A *a, int value) {
	a->value = value;
}
*/
import "C"
import (
	"fmt"
)

func main() {
	var (
		stSize = C.sizeof_struct_A  // 構造体のサイズを取得し (C.sizeof_struct_(構造体名)で取得出来る)
		cSize  = C.size_t(stSize)   // size_tに変換して
		ptr    = C.malloc(cSize)    // malloc(*1)でメモリ確保する。この段階では (void *) なので 
		ptrA   = (*C.struct_A)(ptr) // 適切なポインタ型にキャスト
	)
	defer C.free(ptr) // メモリ動的確保しているので忘れずにfreeすること
	// (*1) ちなみに C.malloc() は、cgo側でカスタムな実装となっておりNULLが返ってくることは絶対にない。失敗時はpanicする。

	C.A_setvalue(ptrA, 999)
	fmt.Printf("value=%d\n", ptrA.value)

	ptrA.value = C.int(111)
	fmt.Printf("value=%d\n", ptrA.value)
}
