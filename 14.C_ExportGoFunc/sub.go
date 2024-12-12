package main

/*
#include <stdio.h>
#include <stdlib.h>
#include "sample.h"
*/
import "C"
import (
	"log"
	"unsafe"
)

//export go_start
func go_start() {
	log.Println("[Go][go_start] start")
}

//export go_end
func go_end() {
	log.Println("[Go][go_end  ] end")
}

//export go_main
func go_main(cId C.int, cData unsafe.Pointer, cLength C.size_t) {
	var (
		id     = int(cId)
		data   = C.GoString((*C.char)(cData))
		length = int(cLength)
	)

	log.Println("[Go][go_main ] called")
	log.Printf("[Go][go_main ] id=%d, data=%q, length=%d\n", id, data, length)
}
