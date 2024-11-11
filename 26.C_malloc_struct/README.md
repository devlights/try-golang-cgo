# これは何？

cgoにてGo側で ```C.malloc()``` を利用して、構造体のメモリを動的確保するサンプルです。

```malloc()``` には、確保するサイズを指定する必要がありますが、通常C言語側では

```c
sizeof(struct A)
```

としますが、cgo側では以下のように構造体のサイズを取得することが出来ます。

```go
var sz = C.sizeof_struct_(構造体)
```

typedefしていない構造体を扱う場合の宣言方法は [25.C_struct_not_typedef](../25.C_struct_not_typedef/) を参照ください。
