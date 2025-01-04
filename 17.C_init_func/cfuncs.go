package main

/*
#include <stdio.h>
#include <stdlib.h>

void my_printf(const char *msg) {
	printf("%s\n", msg);
}
*/
import "C"

func init() {
	println(">> cfuncs.init() called")
	m["cfuncs"] = true
}
