# これは何？

cgoを利用して共有ファイルを生成した場合でも ```init()``` は、ちゃんと呼び出されるのかについてのサンプルです。

結論としては、「ちゃんと呼び出されている」ですが、以下の挙動となります。

- [C言語ネイティブで定義したmy_printf](./cfuncs.go)が初回で呼ばれるときは ```init()``` が呼び出されない
- [CGOで定義したgofunc1](./exports.go)が初回で呼ばれるときに ```init()``` が呼び出される

つまり、CGOの境界に入ったときに ```init()``` が呼び出されるという挙動になります。

```sh
$ task
task: [default] go build -o libgo.so -buildmode=c-shared .
task: [default] gcc -c -o main.o main.c
task: [default] gcc -o capp main.o -L. -lgo
task: [default] ldd ./capp
        linux-vdso.so.1 (0x00007fff5778a000)
        libgo.so => ./libgo.so (0x00007ff4bda66000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007ff4bd836000)
        /lib64/ld-linux-x86-64.so.2 (0x00007ff4bdc1f000)
task: [default] ./capp
[C ] hello world
>> main.init() called
>> cfuncs.init() called
>> exports.init() called
[GO] gofunc1
[GO] map[cfuncs:true exports:true main:true]
```
