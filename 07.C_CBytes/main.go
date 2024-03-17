package main

/*
#cgo CFLAGS: -Wall

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SKIP_BYTES 6

void p1(const void *data) {
	printf("%s\n", (char *)data);
}

void p2(void *data) {
	void *buf = data + SKIP_BYTES;
	printf("%s\n", (char *)buf);
}

void p3(char *data) {
	char *buf = data + SKIP_BYTES;
	printf("%s\n", buf);
}
*/
import "C"

func main() {
	//
	// []byteをC側のバイト列にするには C.CBytes()を利用する.
	//
	var (
		goBytes = []byte("hello Go world")
		cBytes  = C.CBytes(goBytes)
	)

	// C.CBytes() は、malloc() にてメモリを確保しているので、C.free() の呼び出しが必須
	defer C.free(cBytes)

	C.p1(cBytes)
	C.p2(cBytes)
	C.p3((*C.char)(cBytes))
}
