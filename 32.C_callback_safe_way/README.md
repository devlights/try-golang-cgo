# cgo コールバック (関数ポインタ型の利用)

## 概要

本サンプルは、`31.C_callback` の補足となるサンプルです。

`31.C_callback` では、Goの関数ポインタをCに渡す際の一般的なイディオムとして、`(*[0]byte)` への直接キャストを利用しました。本サンプルでは、C側で関数ポインタの `typedef` を定義し、それを利用することで、よりcgoの型変換の仕組みに沿った形で関数ポインタを取得する方法を示します。

最終的にCの関数に渡される値は `31.C_callback` と同じですが、そこに至るまでの過程が異なります。

## `31.C_callback` との違い

一番の違いは、Cのコードブロック内で関数ポインタの型を `typedef` で定義している点です。

```c
typedef int (*callback_fn)(int, int);
```

cgoは、このように `typedef` で定義された関数ポインタ型 `callback_fn` を、Go側では `*[0]byte` 型として特別に扱います。そして、`C.callback_fn` というGoの型としても利用できるようになります。

## コードのポイント

`main.go` では、以下の手順で関数ポインタを取得しています。

1.  `//export` されたGoの関数 `export_func` は、cgoによって `unsafe.Pointer` 型の `C.export_func` という変数として提供されます。
    ```go
    var fn unsafe.Pointer = C.export_func
    ```

2.  Cで `typedef` した `callback_fn` を使い、`unsafe.Pointer` 型の `fn` をキャストします。cgoのルールにより、`C.callback_fn` は `*[0]byte` 型として解釈されるため、結果として `p` は `*[0]byte` 型のポインタになります。
    ```go
    var p *[0]byte = C.callback_fn(fn)
    ```

3.  この `p` は、`31.C_callback` で見た `(*[0]byte)(C.export_func)` という直接的なキャストと全く同じ値になります。
    ```go
    var p2 = (*[0]byte)(C.export_func) // 上の過程を一行で書いたものと同じ
    ```

このように、`typedef` を経由する方法は、cgoが内部で行う型変換をより明示的にコードに表現する方法と言えます。どちらの方法も有効ですが、このサンプルはcgoの挙動をより深く理解する助けとなります。

## 実行方法

```bash
go run .
```

## 実行結果

```
-------------------------------
[C ] x=2, y=3
[Go] sleep 1sec
[Go] x=2, y=3, ans=6
[C ] ans=6
-------------------------------
[C ] x=2, y=3
[Go] sleep 1sec
[Go] x=2, y=3, ans=6
[C ] ans=6
[MAIN] 6:6
```
