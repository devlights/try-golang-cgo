package main

import "log"

var (
	m = make(map[string]bool)
)

func init() {
	println(">> main.init() called")
	log.SetFlags(0)
	m["main"] = true
}

func main() {}
