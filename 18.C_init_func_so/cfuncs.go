package main

/*
#include <stdio.h>

void my_printf(const char *msg) {
	printf("[C ] %s\n", msg);
}
*/
import "C"

func init() {
	m["cfuncs"] = true
}
