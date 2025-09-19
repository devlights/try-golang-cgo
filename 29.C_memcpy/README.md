# これは何？

cgoでmemcpyを利用するサンプルです。

以下のサンプルにもなっています。

```sh
$ task
task: [default] rm -f app liba.*
task: [default] go build -o liba.a -buildmode=c-archive main.go
task: [default] gcc -Wall -o app app.c liba.a
task: [default] ./app
[before] helloworld:
[after ] helloworld:HELLOWORLD
```
