# これは何？

cgoを利用して共有ファイルを生成した場合でも ```init()``` は、ちゃんと呼び出されるのかについてのサンプルです。

結論としては、「ちゃんと呼び出されている」でした。

```sh
$ task
task: [default] go build -o libgo.so -buildmode=c-shared .
task: [default] gcc -c -o main.o main.c
task: [default] gcc -o capp main.o -L. -lgo
task: [default] ldd ./capp
        linux-vdso.so.1 (0x00007ffcc2745000)
        libgo.so => ./libgo.so (0x00007fd9257e0000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fd9255b0000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fd9259ae000)
task: [default] ./capp
[C ] hello world
[GO] gofunc1
[GO] map[cfuncs:true exports:true main:true]
```
