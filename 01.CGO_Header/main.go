package main

/*
#cgo CFLAGS: -I.

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sample.h"

void myPrintf(const char* msg) {
	printf("[C] %s\n", msg);
}

void printStA(const ST_A* st) {
	printf("[C] %d\t%s\n", st->Id, st->Name);
}
*/
import "C" // cgoを利用する場合は必ずこれが必要。この上にコメントでcgoヘッダを記載する。
import (
	"fmt"
	"unsafe"
)

func main() {
	// cgoヘッダで include しておけば、C側の定数を C.定数名 で利用できる
	fmt.Println("[GO]", C.EXIT_SUCCESS, C.EXIT_FAILURE)

	// これはGoでの普通の文字列
	str := "hello world"

	// C側の文字列を利用する場合は C.CString を使う
	// C.CStringで作成した値は (char *) と同じ。
	//
	// 以下のように記載されている
	//      CString converts the Go string s to a C string.
	//      The C string is allocated in the C heap using malloc.
	//      It is the caller's responsibility to arrange for it to be freed, such as by calling C.free
	//      (be sure to include stdlib.h if C.free is needed).
	cStr := C.CString(str)

	// C.CString は、malloc() にてメモリが確保されているので必ず free() が必要
	// C.free() には、(void *) を指定しないといけないので unsafe.Pointer() で (void *) にして渡す
	defer C.free(unsafe.Pointer(cStr))

	// cgoには以下の制約事項がある
	//
	//   - 可変長引数を持つ関数は呼び出せない
	//   - マクロは利用できない
	//
	// printf()は、可変長引数を持っているので cgo から直接呼び出すことは出来ない.
	// 同様に、マクロも cgo からは利用できない。
	C.myPrintf(cStr)

	// C.CStringはC側の文字列（自動的に末尾に 終端記号('\0') が追加される）されたもの。
	// これをGo側の文字列にするには C.GoString を使う。
	//
	// 以下のように記載されている
	//     GoString converts the C string p into a Go string.
	goStr := C.GoString(cStr)

	fmt.Printf("[C.CString ] %[1]v (%[1]T)\n", cStr)
	fmt.Printf("[C.GoString] %[1]v (%[1]T)\n", goStr)

	// 独自定義している構造体もCGOヘッダにてincludeしていれば
	// C.構造体名 で利用できるようになる
	//
	// ただ、文字列が絡むと unsafe.Pointer() の嵐となる
	var stA C.ST_A

	stA.Id = C.int(999)
	C.strncpy((*C.char)(unsafe.Pointer(&stA.Name)), (*C.char)(unsafe.Pointer(cStr)), C.ulong(len(str)+1))

	C.printStA(&stA)

	// stA は、Go側で変数宣言してスタックに確保しているため C.free() は必要ない
	// C.CString は、C側で malloc でヒープメモリに確保されるため C.free() が必要
}
