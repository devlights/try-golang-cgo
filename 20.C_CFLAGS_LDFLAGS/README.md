# これは何？

CGOヘッダーには、C言語で利用する ```CFLAGS```, ```LDFLAGS``` が利用できます。

外部ライブラリを利用する際は、CGOヘッダーにて指定します。

```go
/*
#cgo CFLAGS: -I../lib
#cgo LDFLAGS: -L../lib -lmylib

#include <stdlib.h>
#include "lib.h"
*/
import "C"
```

本サンプルを実行すると以下のようになります。

```sh
$ task
task: [build-lib] make build
gcc -g -O0 -Wall -Wextra -std=c17 -c /workspace/try-golang-cgo/20.C_CFLAGS_LDFLAGS/lib/lib.c -o /workspace/try-golang-cgo/20.C_CFLAGS_LDFLAGS/lib/lib.o 
gcc -g -O0 -Wall -Wextra -std=c17 -fPIC -shared  -o ./libmylib.so /workspace/try-golang-cgo/20.C_CFLAGS_LDFLAGS/lib/lib.o 
task: [build-capp] make build
gcc -g -O0 -Wall -Wextra -std=c17 -c /workspace/try-golang-cgo/20.C_CFLAGS_LDFLAGS/capp/main.c -o /workspace/try-golang-cgo/20.C_CFLAGS_LDFLAGS/capp/main.o 
gcc -g -O0 -Wall -Wextra -std=c17 -L../lib -o ./cApp /workspace/try-golang-cgo/20.C_CFLAGS_LDFLAGS/capp/main.o -lmylib
task: [build-goapp] go build -o goApp
task: [run-capp] LD_LIBRARY_PATH=../lib ./cApp
[C ] hello world
task: [run-goapp] LD_LIBRARY_PATH=../lib ./goApp
[C ] hello world
task: [list-symbols] nm --extern-only capp/cApp   | grep -E ' [T|U] (main|myPrint)'
0000000000001149 T main
                 U myPrint
task: [list-symbols] nm --extern-only goapp/goApp | grep -E ' [T|U] (main|myPrint)'
00000000004642a0 T main
                 U myPrint
task: [ldd] LD_LIBRARY_PATH=lib ldd capp/cApp
        linux-vdso.so.1 (0x00007ffffa77c000)
        libmylib.so => lib/libmylib.so (0x00007f5b71474000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f5b71244000)
        /lib64/ld-linux-x86-64.so.2 (0x00007f5b71480000)
task: [ldd] LD_LIBRARY_PATH=lib ldd goapp/goApp
        linux-vdso.so.1 (0x00007fffd0998000)
        libmylib.so => lib/libmylib.so (0x00007f776fe38000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f776fc08000)
        /lib64/ld-linux-x86-64.so.2 (0x00007f776fe3f000)
```
