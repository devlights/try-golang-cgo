package main

/*
#include <stdlib.h>
*/
import "C"
import "log"

func init() {
	println(">> exports.init() called")
	m["exports"] = true
}

//export gofunc1
func gofunc1() {
	log.Printf("[GO] gofunc1")
	log.Printf("[GO] %v", m)
}
