# これは何？

cgoでは、cgoヘッダーのcgoディレクティブにて ```${SRCDIR}``` という変数が利用できます。

```go
/*
#cgo LDFLAGS: -L${SRCDIR}/libs -llibs
*/
import "C"
```

```#include```の部分には利用出来ません。つまり以下は駄目です。

```go
/*
// これはオッケイ
#cgo LDFLAGS: -L${SRCDIR}/libs -llibs

// 以下は駄目
#include "${SRCDIR}/libs/libs.h"
*/
import "C"
```

この場合は、以下のようにします。

```go
/*
#cgo CFLAGS:  -I${SRCDIR}/libs
#cgo LDFLAGS: -L${SRCDIR}/libs -llibs

#include "libs.h"
*/
import "C"
```


この値は、実行時に現在のソースファイルのディレクトリに置換されます。ディレクトリパスは絶対パスで展開されます。

https://pkg.go.dev/cmd/cgo に以下の記載で記されています。

> When the cgo directives are parsed, any occurrence of the string ${SRCDIR} will be replaced by the absolute path to the directory containing the source file. 

> (cgoディレクティブが解析されるとき、${SRCDIR}という文字列は、ソースファイルを含むディレクトリへの絶対パスに置き換えられます。)

本サンプルを実行すると、以下のようになります。

```sh
$ task
task: [build-lib] make build
gcc -g -O0 -Wall -Wextra -std=c17 -c /workspace/try-golang-cgo/22.C_SRCDIR/libs/libs.c -o /workspace/try-golang-cgo/22.C_SRCDIR/libs/libs.o 
gcc -g -O0 -Wall -Wextra -std=c17 -fPIC -shared  -o ./liblibs.so /workspace/try-golang-cgo/22.C_SRCDIR/libs/libs.o 
task: [build-goapp] go build -o app
task: [run-goapp] ./app
30
```
