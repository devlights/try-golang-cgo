package main

import "C"
import (
	"fmt"
	"unsafe"
)

//export goCallback
func goCallback(cId C.int, cData unsafe.Pointer, cLength C.size_t) {
	var (
		id     = int(cId)
		data   = C.GoString((*C.char)(cData))
		length = int(cLength)
	)

	fmt.Println("[Go][goCallback] called")
	fmt.Printf("[Go][goCallback] id=%d, data=%q, length=%d\n", id, data, length)
}
