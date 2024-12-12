# これは何？

Go側で動的共有ライブラリ(soファイル)を作成して、C言語側からexportされた関数を呼び出すサンプルです。

Goでsoファイルを作成するには、 ```go build``` 時に ```-buildmode=c-shared``` オプションを付与します。

```-buildmode=c-shared``` を付与してビルドするとsoファイルとヘッダファイルが生成されます。(exportしている関数がある場合)

本サンプルを実行すると以下のようになります。

```sh
$ task
task: [build-sofile] go build -o libapp.so -buildmode=c-shared main.go
task: [build-c] gcc -o app app.c -I. -L. -lapp
task: [run] ./app
30
```

## 参考情報

- [Goメモ-356 (Go側からsoファイルを作成してPythonとCで利用)(c-shared)](https://devlights.hatenablog.com/entry/2023/11/30/073000)
