package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>

// open(2) は glibc で varargs 宣言されている。そのため cgo から直接呼べないので
// 固定引数の static inline ラッパーを用意して回避する。
// macOSの場合は、openが固定引数で見えるので、この問題は発生しない。
static inline int go_open(const char *p, int flags) {
	return open(p, flags);
}
*/
import "C"
import (
	"context"
	"errors"
	"fmt"
	"log"
	"slices"
	"syscall"
	"unsafe"
)

const (
	exists    = "README.md"
	notExists = "NOT-EXISTS.md"
)

func main() {
	log.SetFlags(0)

	var (
		ctx = context.Background()
		err error
	)
	if err = run(ctx); err != nil {
		log.Println(err)
	}
}

func run(_ context.Context) error {
	var (
		p1 = C.CString(exists)
		p2 = C.CString(notExists)
	)
	defer C.free(unsafe.Pointer(p1))
	defer C.free(unsafe.Pointer(p2))

	type (
		loopVal struct {
			fpath string
			cpath *C.char
		}
	)
	var (
		fd   C.int
		err  error
		vals = []loopVal{{exists, p1}, {notExists, p2}}
	)
	for v := range slices.Values(vals) {
		// cgo経由でCの関数を呼び出す場合、Goの多値返却の仕組みにより
		// 本来の戻り値にプラスerrorが返ってくる。この error には、C側のerrnoの値がセットされる。
		//
		// errも返ってきているが C の関数を扱う場合は err != nil だけで判定するのは危険。
		// Cのセオリー通り、まず戻り値で判定し、その後で必要であれば err を使う。
		// errが存在する場合 syscall.Errno となる。（呼び出しが成功している場合、通常nilとなる)
		fd, err = C.go_open(v.cpath, C.int(syscall.O_RDONLY))
		if fd == -1 {
			// err は、実際には syscall.Errno となっている
			// 今回の場合では、ファイルが存在しないため、syscall.ENOENT となる。
			var (
				errno   = err.(syscall.Errno)
				eno     = uintptr(errno)
				isNOENT = errors.Is(err, syscall.ENOENT)
			)
			return fmt.Errorf("NG: open(%s): fd=%d (%w)(0x%x:ENOENT=%v)", v.fpath, fd, err, eno, isNOENT)
		}

		C.close(fd)

		log.Printf("OK: open(%s): fd=%d (%v)", v.fpath, fd, err)
	}

	return nil
}
