package main

/*
#include <string.h>
#include <stdlib.h>

typedef struct _Data {
	char name[10];
} Data;

static int Init_Data(Data* v) {
	const char msg[] = "Hello Go";

	strncpy(v->name, msg, sizeof(msg)-1);
	v->name[sizeof(msg)-1] = '\0';

	return EXIT_SUCCESS;
}
*/
import "C"
import (
	"log"
	"unsafe"
)

func init() {
	log.SetFlags(0)
}

func main() {
	var data C.Data
	C.Init_Data(&data)

	/*
		Data構造体は以下のような形でcgo側に展開される

		type _Ctype_struct___0 struct {
			name [10]_Ctype_char
		}

		Goで配列は要素数も含めて型となるため
		そのままでは (char *) としては利用できない。

		Cの文字列をGoの文字列として利用する際に
		使う C.GoString() は (char *) を引数に要求する。

		変換を行うには一旦 unsafe.Pointer にしてから
		望みの型へとキャストする。

		尚、Goの世界ではC言語のように「配列の先頭要素は暗黙的に配列自体のポインタを表す」と
		いうルールにはならないので、明示的に先頭要素のポインタを指定する。
	*/
	p := unsafe.Pointer(&data.name[0]) // unsafe.Pointerにしてから
	c := (*C.char)(p)                  // (char *) にキャストし
	s := C.GoString(c)                 // Goの文字列に変換

	log.Println(s)
}
