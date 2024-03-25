# これは何？

cgoとdlopen関数を使って、既に存在している共有ライブラリ内の既存関数と全く同じ関数書式を持つ関数をcgo側で定義して、呼び出しをフックするサンプルです。

実行すると以下のようになります。

```sh
$ task
task: [build-clib] gcc -fPIC -shared -o libclib.so lib.c
task: [build-golib] go build -o libgolib.so -buildmode=c-shared *.go
task: [build-cprg] gcc -c -o main.o main.c
task: [build-cprg] gcc -o capp main.o -L. -lgolib
task: [default] LD_LIBRARY_PATH=. ./capp
2024/03/25 02:36:06 [Go][インターセプト][my_strlen] result=10
my_strlen=10
task: [build-cprg-original] gcc -o capp main.o -Lclib -lclib
task: [default] LD_LIBRARY_PATH=clib ./capp
my_strlen=10
```
