# C.ExportGoFunc

```//export XXXX``` とすることで、Goの関数をC側に公開することが出来る。

詳細は、[cmd/cgo](https://pkg.go.dev/cmd/cgo#hdr-C_references_to_Go)にて記載されている。

本サンプルでは、C側にて予め定義されている処理のメイン処理がテンプレートパターンのように決まった名前の関数を順次コールしていくという流れになっている。

コールされる側の関数のプロトタイプは以下。```sample.h```に定義している。

```c
extern void go_start();
extern void go_end();
extern void go_main(int id, void *data, size_t length);
```

C側のメイン処理は以下のような形。以下がGoのmain()よりcgo経由で呼び出される。

```c
void c_run() {
	go_start();

	int id = 100;
	char data[] = "hello world";
	size_t szData = strnlen(data, 20);

	go_main(id, data, szData);

	go_end();
}
```

```go_xxxx``` となっている部分がGo側で定義し、exportしているものとなる。関数定義は ```sub.go``` にある。

ビルドすると以下のように、ちゃんとそのままの名前でシンボルテーブルに登録されている。

```sh
$ task build
task: [build] go build -o app
task: [build] nm ./app | grep -E "T (go_start|go_end|go_main)$"
000000000048d3b0 T go_end
000000000048d420 T go_main
000000000048d340 T go_start
```

実行すると以下のようになる。

```sh
$ task
task: [default] go run .
2024/03/18 10:48:22 [Go][go_start] start
2024/03/18 10:48:22 [Go][go_main ] called
2024/03/18 10:48:22 [Go][go_main ] id=100, data="hello world", length=11
2024/03/18 10:48:22 [Go][go_end  ] end
```

---

補足情報として nm コマンドにて ```T``` となっているものは以下の意味を持つ。

> nmコマンドの結果におけるTは、そのシンボルがテキストセクション（テキストセグメント）に存在し、外部リンク可能（つまり、他のオブジェクトファイルや実行可能ファイルから参照できる）状態であることを表しています。

> T : "Text"セクションにある非静的（グローバル）関数または変数を表します。これは、実行可能コードが含まれており、外部からリンケージ（参照や呼び出しが可能）であることを意味します。要するに、Tでマークされたシンボルは、グローバルに公開されている関数や変数であり、プログラム内で定義されている関数や変数のアドレスを示します。これらはプログラムの実行コード部分に存在し、リンカによって他のオブジェクトファイルと結合することが可能です。
