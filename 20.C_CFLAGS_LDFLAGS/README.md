# これは何？

CGOヘッダーには、C言語で利用する ```CFLAGS```, ```LDFLAGS``` が利用できます。

外部ライブラリを利用する際は、CGOヘッダーにて指定します。

```go
/*
#cgo CFLAGS: -I../lib
#cgo LDFLAGS: -L../lib -lmylib

#include <stdlib.h>
#include "lib.c"
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
gcc -g -O0 -Wall -Wextra -std=c17 -c /workspace/try-golang-cgo/20.C_CFLAGS_LDFLAGS/capp/main.c -o /workspace/try-golang-cgo/20.C_CFLAGS_LDFLAGS/capp/main.o -I../lib
gcc -g -O0 -Wall -Wextra -std=c17 -L../lib -o ./cApp /workspace/try-golang-cgo/20.C_CFLAGS_LDFLAGS/capp/main.o -lmylib
task: [build-goapp] go build -o goApp
task: [run-capp] LD_LIBRARY_PATH=../lib ./cApp
[C ] hello world
task: [run-goapp] ./goApp
[C ] hello world
task: [list-symbols] nm --extern-only capp/cApp   | grep -E ' [T|U] (main|myPrint)'
0000000000001149 T main
                 U myPrint
task: [list-symbols] nm --extern-only goapp/goApp | grep -E ' [T|U] (main|myPrint)'
000000000045c600 T main
0000000000461030 T myPrint
task: [ldd] LD_LIBRARY_PATH=lib ldd capp/cApp
        linux-vdso.so.1 (0x00007ffe22799000)
        libmylib.so => lib/libmylib.so (0x00007f29c42ce000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f29c409e000)
        /lib64/ld-linux-x86-64.so.2 (0x00007f29c42da000)
task: [ldd] ldd goapp/goApp
        linux-vdso.so.1 (0x00007ffcd01a7000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fcf8ac0d000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fcf8ae3f000)
```

よく見ると、C言語のアプリの方は nm, ldd の結果ともに外部ライブラリの参照となっているが

cgoの方は、soファイルが取り込まれているのか nm コマンドにて myPrint 関数が ```T``` で表示される。(C言語側は ```U```)

また、ldd で見た場合に、```libmylib.so``` が出てこない。
