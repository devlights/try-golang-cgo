package main

import "log"

var (
	m = make(map[string]bool)
)

func init() {
	log.SetFlags(0)
	m["main"] = true
}

func main() {}
