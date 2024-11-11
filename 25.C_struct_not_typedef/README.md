# これは何？

cgoにてC側で構造体を定義して利用する場合

```c
struct A { ... };
```

と

```c
typedef struct A TA;
```

では、Go側での宣言で指定する型が異なるというサンプルです。

```struct A { ... };``` とした場合は

```go
var a C.struct_(構造体名)
```

と宣言します。

```typedef struct A TA;``` とした場合は

```go
var ta C.TA
```

と宣言します。
