package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
)

func main() {
	//
	// C言語側の型を利用するには C.型名 とする。
	//
	goint := 10
	cint := C.int(goint)

	fmt.Printf("Go=%d\tC=%d\n", goint, cint)

	gobyte := byte(127)
	cbyte := C.char(gobyte)

	fmt.Printf("Go=%d\tC=%d\n", gobyte, cbyte)
}
