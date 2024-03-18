# C.malloc

```C.malloc()```は、C言語側の```malloc()```を呼び出すが、特殊仕様があり

実際にはGo側のラッパー関数が呼び出される。このラッパー関数は決してnilを返すことはない。

メモリ確保に失敗した場合は、パニックになっている。

以下、ドキュメント記載。(https://pkg.go.dev/cmd/cgo)

> As a special case, C.malloc does not call the C library malloc directly but instead calls a Go helper function that wraps the C library malloc but guarantees never to return nil. If C's malloc indicates out of memory, the helper function crashes the program, like when Go itself runs out of memory. Because C.malloc cannot fail, it has no two-result form that returns errno.

> 特殊なケースとして、C.mallocはCライブラリのmallocを直接呼び出すのではなく、CライブラリのmallocをラップするGoヘルパー関数を呼び出しますが、決してnilを返さないことが保証されています。Cのmallocがメモリ不足を示すと、Go自身がメモリ不足になったときのように、ヘルパー関数がプログラムをクラッシュさせる。C.mallocは失敗しないので、errnoを返す2つの結果形式はありません。
