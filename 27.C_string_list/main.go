package main

import "C"
import (
	"fmt"
	"unsafe"
)

//export gofunc
func gofunc(argc C.int, argv **C.char) {
	//
	// Cの **char をGoのスライスに変換するには
	//   unsafe.Slice()
	// が使える。(**char --> []*C.char)
	//
	// REF: https://go.dev/wiki/cgo#turning-c-arrays-into-go-slices
	//
	var (
		length = int(argc)
		slice  []*C.char
	)
	slice = unsafe.Slice(argv, length)

	//
	// 後は要素ごとに C.GoString() でGoの文字列に変換出来る
	// (*C.char --> string)
	//
	for _, v := range slice {
		fmt.Println(C.GoString(v))
	}
}

func main() {}
