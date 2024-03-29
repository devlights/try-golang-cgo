package main

/*
#include <stdio.h>
#include <stdlib.h>

typedef struct {
	int v;
	char v2[20];
} ST_Y;

typedef struct {
	ST_Y y;
} ST_X;
*/
import "C"
import (
	"fmt"
)

func main() {
	//
	// 特例として、Go側にてC言語側の構造体を変数宣言した場合
	// 自動的にゼロクリアされた状態となる。（memset(&stx, 0, sizeof(stx))と同じ状態)
	// これは、Goの変数初期化のルールが適用されるため（ゼロ値）
	//
	var stx C.ST_X
	stx.y.v = 100

	fmt.Printf("%+v\n", stx)
}
