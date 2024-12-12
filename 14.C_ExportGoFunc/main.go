package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sample.h"

void c_run() {
	go_start();

	int id = 100;
	char data[] = "hello world";
	size_t szData = strnlen(data, 20);

	go_main(id, data, szData);

	go_end();
}
*/
import "C"

func main() {
	C.c_run()
}
