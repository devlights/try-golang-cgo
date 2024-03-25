# これは何？

cgoを利用している場合でも ```init()``` は、ちゃんと呼び出されるのかについてのサンプルです。

結論としては、「ちゃんと呼び出されている」でした。

```sh
$ task
task: [default] go run .
map[cfuncs:true exports:true gofuncs:true main:true]
helloworld
func1
```
