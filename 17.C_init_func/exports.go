package main

import "C"
import "log"

func init() {
	m["exports"] = true
}

//export func1
func func1() {
	log.Println("func1")
}
