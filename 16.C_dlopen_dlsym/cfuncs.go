package main

/*
size_t call_my_strlen(void *fn, const char *s) {
	size_t (*my_strlen_fn)(const char *);

	my_strlen_fn = (size_t (*)(const char *))fn;
	return my_strlen_fn(s);
}
*/
import "C"
