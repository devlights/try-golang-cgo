package main

/*
#cgo CFLAGS: -Wall -Wextra -g3 -O0

#include <stdio.h>

// Go側でexportした関数のプロトタイプ宣言
extern int export_func(int x, int y);

// 実際に呼び出すC側の関数。引数に関数ポインタを要求している。
int c_func(int x, int y, int (*fn)(int, int)) {
	printf("[C ] x=%d, y=%d\n", x, y);
	int ans = fn(x, y);
	printf("[C ] ans=%d\n", ans);

	return ans;
}
*/
import "C"
import "fmt"

func main() {
	var (
		fn = (*[0]byte)(C.export_func) // 関数の先頭ポインタを渡す必要があるためbyteのゼロバイト目のポインタとする
		x  = C.int(2)
		y  = C.int(3)
		z  = C.c_func(x, y, fn)
	)
	fmt.Printf("[Go] z=%d\n", int(z))
}
