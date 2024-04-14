# これは何？

cgoにて利用するコンパイラを変更するには環境変数 ```CC``` で設定します。

(cgoヘッダで対応しているのは CFLAGS, CPPFLAGS, CXXFLAGS, FFLAGS と LDFLAGS です ([Using cgo with the go command](https://pkg.go.dev/cmd/cgo#hdr-Using_cgo_with_the_go_command)))

```sh
$ CC=gcc go run main.go
$ CC=clang go run main.go
```

ちゃんと指定されているかどうかは、```go build -x``` した際に出力される結果の ```-extld=xxx``` の部分を見ると分かります。

```sh
$ CC=clang go run -x main.go 2>&1 | grep -o '\-extld=[^ ]*'
-extld=clang
```
