package main

/*
#cgo CFLAGS: -g3 -O0 -Wall -Wextra

#include <stdio.h>
*/
import "C"

//export c_func
func c_func(x, y C.int) C.int {
	return x * y
}

func main() {}
