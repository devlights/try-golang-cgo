//go:build run

package main

/*
#cgo LDFLAGS: -ldl

#include <dlfcn.h>
#include <stdlib.h>

typedef size_t (*strlen_t)(const char *);

// dlopenでlibc.so.6を開き、strlen関数のアドレスを取得する
strlen_t getStrlenFunc() {
    void* handle = dlopen("libc.so.6", RTLD_LAZY);
    if (!handle) {
        return NULL;
    }
    dlerror(); // dlerrorの状態をクリア

    strlen_t func = (strlen_t)dlsym(handle, "strlen");
    if (dlerror() != NULL) {
        return NULL;
    }

    return func;
}

size_t callstrlen(const char *s) {
	strlen_t fn = getStrlenFunc();
	return fn(s);
}
*/
import "C"
