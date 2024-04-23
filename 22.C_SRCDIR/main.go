package main

/*
#cgo CFLAGS:  -I${SRCDIR}/libs
#cgo LDFLAGS: -L${SRCDIR}/libs -llibs

#include "libs.h"
*/
import "C"
import "log"

func init() {
	log.SetFlags(0)
}

func main() {
	var (
		x = C.int(10)
		y = C.int(20)
	)

	log.Println(int(C.add(x, y)))
}
