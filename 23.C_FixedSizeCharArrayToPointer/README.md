# これは何？

C側の構造体にて固定要素数の文字配列が定義されている場合に、それをGo側で文字列に変換する方法についてのサンプルです。

```c
typedef struct _Data {
	char name[10];
} Data;
```

上記のC構造体は以下のような形でcgo側に展開されます。

```go
type _Ctype_struct__Data struct {
    name [10]_Ctype_char
}
```

Goで配列は要素数も含めて型となるため、そのままでは (char *) としては利用できません。

Cの文字列をGoの文字列として利用する際に使う C.GoString() は (char *) を引数に要求します。

変換を行うには一旦 unsafe.Pointer にしてから、望みの型へとキャストします。

尚、Goの世界ではC言語のように「配列の先頭要素は暗黙的に配列自体のポインタを表す」というルールにはなりませんので、明示的に先頭要素のポインタを指定します。
