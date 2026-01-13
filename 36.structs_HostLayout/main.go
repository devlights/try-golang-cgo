package main

/*
#cgo CFLAGS: -std=c99

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

typedef struct {
	uint8_t val1; // 1byte
	              // パディング 3byte
	int32_t val2; // 4byte
	              // パディング 0byte (8バイト境界になるため)
	int64_t val3; // 8byte
} c_struct;

c_struct* new_c_struct() {
	c_struct *v = (c_struct *)malloc(sizeof(c_struct));
	if (v == NULL) {
		return NULL;
	}

	v->val1 = 1;
	v->val2 = 2;
	v->val3 = 3;

	return v;
}

void free_c_struct(c_struct *v) {
	free(v);
}
*/
import "C"
import (
	"fmt"
	"structs"
	"unsafe"
)

type (
	Go_Struct struct {
		_    structs.HostLayout // 「ホストプラットフォームのCコンパイラと同じルールでレイアウトせよ」と指示
		val1 uint8              // 1byte
		val2 int32              // 4byte
		val3 int64              // 8byte
	}
)

func main() {
	//
	// Cと同じ形の構造体をGo側にも定義し、その構造体にstructs.HostLayoutを適用しておき
	// C側とGo側で全く同じように見えるかどうかを確認する。
	//
	// structs.HostLayoutは、Go1.23で追加された。
	//
	var (
		cPtr  *C.c_struct
		goPtr *Go_Struct
	)
	cPtr = C.new_c_struct()
	defer C.free_c_struct(cPtr)

	// Goのポインタとしてキャスト
	goPtr = (*Go_Struct)(unsafe.Pointer(cPtr))

	// 値を確認
	fmt.Printf("[C ] val1=%v, val2=%v, val3=%v\n", cPtr.val1, cPtr.val2, cPtr.val3)
	fmt.Printf("[Go] val1=%v, val2=%v, val3=%v\n", goPtr.val1, goPtr.val2, goPtr.val3)

	// パディングの確認
	var (
		addrVal1 = uintptr(unsafe.Pointer(&goPtr.val1))
		addrVal2 = uintptr(unsafe.Pointer(&goPtr.val2))
		addrVal3 = uintptr(unsafe.Pointer(&goPtr.val3))
		val1Size = unsafe.Sizeof(goPtr.val1)
		val2Size = unsafe.Sizeof(goPtr.val2)
		padding1 = addrVal2 - addrVal1 - val1Size
		padding2 = addrVal3 - addrVal2 - val2Size
	)
	fmt.Printf("アドレス：  (1) %#x, (2) %#x, (3) %#x\n", addrVal1, addrVal2, addrVal3)
	fmt.Printf("パディング: (1) %v bytes, (2) %v bytes\n", padding1, padding2)
}
