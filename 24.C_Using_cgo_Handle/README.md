# これは何？

```cgo.Handle``` を用いて、CとGoの間で値をやり取りするサンプルです。

cgo.Handleは、GoとCの間で安全にGoの値を渡すためのメカニズムです。cgo.Handleは整数値(Cでは```uintptr_t```, Goでは```uintptr```)であり、任意のGoの値を表現できます。

このハンドルを使用することで、Goのポインタを直接Cに渡すことなく、GoとCの間でデータをやり取りすることができます。

これにより、cgoのポインタ渡しルールを破ることなく、安全にデータをやり取りできます。

主にCからGoでエクスポートした関数を呼ぶ際などに利用します。


## Cの```uintptr_t```型

uintptr_tは、C言語における特別な整数型で、ポインタの値を格納するために十分な大きさの符号なし整数型です。

この型は、ポインタを整数として扱う際に使用され、特にアドレス演算やシステムコールなどの低レベルの操作で役立ちます。

```stdint.h```に含まれています。

```c
#include <stdio.h>
#include <stdint.h>

int main() {
    int x = 42;
    int *p = &x;
    uintptr_t u = (uintptr_t)p;
    printf("Pointer: %p, Uintptr: %lu\n", (void*)p, (unsigned long)u);
    return 0;
}
```

## Goの```uintptr```型

uintptrは、Go言語における特別な整数型で、ポインタの値を格納するために十分な大きさの符号なし整数です。

これは、メモリアドレスを直接扱うために使用され、特にシステムコールや低レベルのメモリ操作で重要な役割を果たします

```go
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var x int = 42
    p := unsafe.Pointer(&x)
    u := uintptr(p)
    fmt.Printf("Pointer: %p, Uintptr: %d\n", p, u)
}
```

## 参考情報

[cgo.Handle](https://pkg.go.dev/runtime/cgo@go1.22.3#Handle)
