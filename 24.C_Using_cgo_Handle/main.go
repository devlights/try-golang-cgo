package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>
#include <stdint.h>

// Go側でexternされる関数
extern void GoCallback(uintptr_t handle);
extern void GoCallback2(void *context);

// Goのコールバックを呼び出すためのラッパー関数
static void C_Callback(uintptr_t handle) { GoCallback(handle);   }
static void C_Callback2(void *context)   { GoCallback2(context); }
*/
import "C"
import (
	"log"
	"runtime/cgo"
	"unsafe"
)

type GoStruct struct {
	Value1 string
	Value2 string
	Value3 int
}

//export GoCallback
func GoCallback(handle C.uintptr_t) {
	var (
		cgoHandle = cgo.Handle(handle)
		value     = cgoHandle.Value().(GoStruct)
	)
	defer cgoHandle.Delete() // 利用が終わったら廃棄

	log.Printf("[HANDLE] value  ptr=%p", &value)
	log.Printf("[GoCallback ] %v", value)
}

//export GoCallback2
func GoCallback2(context unsafe.Pointer) {
	var (
		cgoHandle = (*(*cgo.Handle)(context)) // *cgo.Handleにキャストして実体取得
		value     = cgoHandle.Value().(*GoStruct)
	)
	defer cgoHandle.Delete() // 利用が終わったら廃棄

	log.Printf("[HANDLE] value2 ptr=%p", value)
	log.Printf("[GoCallback2] %v", value)
}

func init() {
	log.SetFlags(0)
}

func main() {
	// -----------------------------------
	// cgo.Handle を uintptr_t として利用
	//
	// Goの値を cgo.Handle でラップ
	var (
		value  = GoStruct{"hello", "world", 999}
		handle = cgo.NewHandle(value)
		ptr    = C.uintptr_t(handle)
	)

	// Cを経由して、Goの関数に戻り処理する
	log.Printf("[MAIN  ] value  ptr=%p", &value)
	C.C_Callback(ptr)

	log.Println("-----------------------------------")

	// -----------------------------------
	// cgo.Handle を void * として利用
	//
	// Goの値を cgo.Handle でラップ
	var (
		value2  = &GoStruct{"world", "hello", 888}
		handle2 = cgo.NewHandle(value2)
		ptr2    = unsafe.Pointer(&handle2)
	)

	// Cを経由して、Goの関数に戻り処理する
	log.Printf("[MAIN  ] value2 ptr=%p", value2)
	C.C_Callback2(ptr2)

	/*
	   $ task
	   task: [run] go build -o app .
	   task: [run] ./app
	   [MAIN  ] value  ptr=0xc00011e0c0
	   [HANDLE] value  ptr=0xc00011e150
	   [GoCallback ] {hello world 999}
	   -----------------------------------
	   [MAIN  ] value2 ptr=0xc00011e1b0
	   [HANDLE] value2 ptr=0xc00011e1b0
	   [GoCallback2] &{world hello 888}
	*/
}
