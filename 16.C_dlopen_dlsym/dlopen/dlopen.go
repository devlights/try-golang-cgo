package dlopen

/*
#cgo LDFLAGS: -ldl

#include <dlfcn.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

func OpenLib(path string) unsafe.Pointer {
	var (
		cPath    = C.CString(path)
		cPathPtr = unsafe.Pointer(cPath)
		handle   unsafe.Pointer
	)
	defer C.free(cPathPtr)

	handle = C.dlopen(cPath, C.RTLD_LAZY)
	return handle
}

func GetSym(handle unsafe.Pointer, symbol string) unsafe.Pointer {
	var (
		cSymbol    = C.CString(symbol)
		cSymbolPtr = unsafe.Pointer(cSymbol)
		funcPtr    unsafe.Pointer
	)
	defer C.free(cSymbolPtr)

	// 本来、dlsym()を呼び出す前に dlerror() を呼び出し、エラー状態をクリアしてから
	// dlsym()を呼び出す。その後に、dlerror() でエラーを確認するのが正当な流れであるが割愛

	funcPtr = C.dlsym(handle, cSymbol)
	return funcPtr
}

func CloseLib(handle unsafe.Pointer) {
	// 本来、dlclose()を呼び出す前に dlerror() を呼び出し、エラー状態をクリアしてから
	// dlclose()を呼び出す。その後に、dlerror() でエラーを確認するのが正当な流れであるが割愛
	C.dlclose(handle)
}
