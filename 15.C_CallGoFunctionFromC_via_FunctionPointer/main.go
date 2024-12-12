package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sample.h"

// func_with_callback関数にて要求される関数ポインタを型定義
typedef void (*c_callback)(int, void *, size_t);

// sub.go にて export しているGo側の関数のプロトタイプ
// プロトタイプと実装を同じファイルで書き込んではいけない (リンカーにてエラーになる)
void goCallback(int, void *, size_t);

// cgo側より呼び出すラッパー関数
void go_func_with_callback(int id, c_callback cb) {
	func_with_callback(id, cb);
}
*/
import "C"

func main() {
	//
	// cgo経由でコールバックを用意して呼び出し
	//
	var (
		id       = C.int(100)
		callback = C.c_callback(C.goCallback)
	)

	C.go_func_with_callback(id, callback)
}
