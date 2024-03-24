package main

/*
#include <stdlib.h>
#include "clib/lib.h"

size_t call_clib_my_strlen(void *fn, const char *s) {
	size_t (*clib_my_strlen_fn)(const char *);

	clib_my_strlen_fn = (size_t (*)(const char *))fn;
	return clib_my_strlen_fn(s);
}
*/
import "C"
