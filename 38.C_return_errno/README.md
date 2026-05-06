# cgoのC関数呼び出しとエラー値

cgo では C 関数を多値返却形式で呼び出すと、Go コンパイラが呼び出し直後に errno を読み取り、error インターフェースとして返します。

[Go references to C](https://pkg.go.dev/cmd/cgo#hdr-Go_references_to_C) に以下の記載があります。

> Any C function (even void functions) may be called in a multiple assignment context to retrieve both the return value (if any) and the C errno variable as an error (use _ to skip the result value if the function returns void). 

```go
n, err = C.sqrt(-1)
_, err := C.voidFunc()
var n, err = C.sqrt(1)
```

```go
ret, err := C.someFunc(args...)
//          ^^^
//          これが errno を syscall.Errno にラップしたもの
```

内部的には以下の型変換が行われています。

```
errno (C の int)
  → syscall.Errno (Go の uintptr、error インターフェースを実装)
  → error インターフェースとして返される
```

## 重要な注意点

> err != nil だけで判定してはいけない。

C 関数が成功しても errno が前回呼び出しの残留値のままの場合があります。

正しいパターンは「戻り値で成否を判断し、失敗時に err を参照する」です。
