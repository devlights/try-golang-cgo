package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const int BUF_SIZE = 15*sizeof(char);

void *buf;

void myAlloc() {
	// MEMO: エラー処理は割愛
	buf = malloc(BUF_SIZE);
	memset(buf, 0, BUF_SIZE);
}

void myAssign() {
	strncpy(buf, "hello C  World", BUF_SIZE);
}

void myPrint() {
	printf("[C ] %s %d\n", (char *)buf, BUF_SIZE);
}

void myFree() {
	free(buf);
}

*/
import "C"
import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

func main() {
	//
	// C側でメモリを確保しバッファにデータを埋める
	//
	C.myAlloc()
	defer C.myFree()

	C.myAssign()
	C.myPrint()

	//
	// Go側の[]byteとして取得
	//
	var (
		cBuf  unsafe.Pointer = C.buf
		goBuf []byte         = C.GoBytes(cBuf, C.BUF_SIZE)
	)

	fmt.Printf("[GO] %s %d\n", string(goBuf), len(goBuf))
	fmt.Printf("[GO] %v\n", hex.Dump(goBuf))
}
