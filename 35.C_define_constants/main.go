package main

/*
#include <stdio.h>
#include <stdlib.h>

#define VALUE1    ((unsigned long)999)
#define MACRO1(x) (printf("[C] %d\n", x))

static inline void macro1(int x) {
	MACRO1(x);
}
*/
import "C"
import "fmt"

func main() {
	//
	// cgoでは
	//   - defineで定義したマクロ定数は見える
	//   - マクロは見えない
	//
	var (
		v = C.ulong(C.VALUE1)
	)
	fmt.Printf("v=%d\n", v)

	// これは無理。マクロは見えない。
	// C.MACRO1(100)

	// C側でラッパー関数用意して呼び出す
	C.macro1(C.int(888))
}
