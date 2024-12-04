# これは何？

C言語の文字列のリスト、つまり ```**char``` を cgo で扱うサンプルです。

```**char``` を ```[]*C.char``` に変換してから、各要素を ```C.GoString()``` します。

```[]*C.char``` への変換には ```unsafe.Slice()``` が使えます。(Go 1.17以降)

## 参考情報

- [Turning C arrays into Go slices](https://go.dev/wiki/cgo#turning-c-arrays-into-go-slices)
