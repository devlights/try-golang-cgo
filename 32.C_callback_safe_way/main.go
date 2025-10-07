package main

/*
#cgo CFLAGS: -Wall -Wextra -g3 -O0

#include <stdio.h>

// c_func関数が要求している関数ポインタのtypedef定義
typedef int (*callback_fn)(int, int);

// Go側でexportした関数のプロトタイプ宣言
extern int export_func(int x, int y);

// 実際に呼び出すC側の関数。引数に関数ポインタを要求している。
int c_func(int x, int y, int (*fn)(int, int)) {
	printf("-------------------------------\n");
	printf("[C ] x=%d, y=%d\n", x, y);
	int ans = fn(x, y);
	printf("[C ] ans=%d\n", ans);

	return ans;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	//
	// 関数ポインタを取るのに (*[0]byte) にキャストするというやり方が
	// よく利用されているイディオムであるが、安全な方法としては unsafe.Pointer を経由させることである
	//
	// ただし、この場合は 関数ポインタ を表現する typedef が必要となる
	//
	// (*[0]byte) は
	//   - [0]byte が要素ゼロの配列を表し、*[0]byteでそれのポインタとなる
	//   - 要素ゼロの配列へのポインタは、実質的にメモリアドレスそのものを表現する
	//   - Goの型システム上、unsafe.Pointerは任意のポインタ型に変換可能
	// という理屈で変換可能となっている
	//

	var (
		fn unsafe.Pointer = C.export_func             // cgo生成で var export_func unsafe.Pointer となる
		p  *[0]byte       = C.callback_fn(fn)         // cgo生成で type callback_fn *[0]byte となる
		p2                = (*[0]byte)(C.export_func) // 上の過程を省いたもの
	)

	//
	// cgo生成で
	//     func c_func(p0 C.int, p1 C.int, p2 *[0]byte) (r1 C.int)
	// となるため、上の p が渡せるようになる。
	//
	// 結局 *[0]byte を渡せば良いので、cgoでは上の過程を省いて 31.C_callback のように
	//     p = (*[0]byte)(C.export_func)
	// とすることが多い。
	//
	var (
		x = C.int(2)
		y = C.int(3)

		z  = C.c_func(x, y, p)
		z2 = C.c_func(x, y, p2)
	)
	fmt.Printf("[MAIN] %d:%d\n", int(z), int(z2))
}
