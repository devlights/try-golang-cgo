package main

/*
#include <stdio.h>
#include <stdlib.h>

struct A {
	int value;
};

typedef struct A TA;

void p(const struct A *a) {
	printf("value=%d\n", a->value);
}
*/
import "C"

func main() {
	var (
		a  C.struct_A // typedefしていない構造体は C.struct_(構造体名) という型となる
		ta C.TA       // typedefしているものは C.型名 でアクセス可能
	)

	//
	// struct A {}; と定義した場合は以下のような展開となる。
	//		type _Ctype_struct_A struct { ... }
	// typedef struct A TA; と定義した場合は以下のような展開となる。
	//		type _Ctype_TA = _Ctype_struct_A
	//

	a.value = C.int(10)
	ta.value = C.int(20)

	C.p(&a)
	C.p(&ta)
}
