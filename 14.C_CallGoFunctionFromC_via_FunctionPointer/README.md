# C.CallGoFuncFromC_via_FunctionPointer

[cgoのドキュメント](https://pkg.go.dev/cmd/cgo)には以下の記載があり、cgoではCの関数ポインタがサポートされていない。しかし、下記のようにやり取りする方法がある。

> Calling C function pointers is currently not supported, however you can declare Go variables which hold C function pointers and pass them back and forth between Go and C. C code may call function pointers received from Go. 

> Cの関数ポインターの呼び出しは現在サポートされていませんが、Cの関数ポインターを保持するGo変数を宣言し、GoとCの間で受け渡しすることができます。

C側にて関数ポインタを引数に要求する関数に、Go側で定義した関数を設定するのは手間がかかるが以下のようにする。

1. 関数ポインタのtypedefが存在しない場合はC側で定義する
1. Goの関数をexportしてCの世界に公開する
1. exportしたGoの関数をtypedefした定義でラップする
1. ラップした値を引数に受取り、実際のCの関数を呼び出すラッパー関数を用意

詳細については、ソースコードを参照。

なお、注意点としてexportするGoの関数定義とcgoのプロトタイプ宣言を同じファイルでしてはいけない。（リンカーでエラーとなる）

この点については、[cgoのドキュメント](https://pkg.go.dev/cmd/cgo)に以下の記載がある。

> Using //export in a file places a restriction on the preamble: since it is copied into two different C output files, it must not contain any definitions, only declarations. If a file contains both definitions and declarations, then the two output files will produce duplicate symbols and the linker will fail. To avoid this, definitions must be placed in preambles in other files, or in C source files.

> ファイル内で//exportを使用すると、プリアンブルに制約が課される。プリアンブルは2つの異なるC出力ファイルにコピーされるため、定義が含まれてはならず、宣言のみが含まれる。ファイルに定義と宣言の両方が含まれていると、2つの出力ファイルに重複したシンボルが生成され、リンカは失敗する。これを避けるには、定義を他のファイルのプリアンブルか、Cのソース・ファイルに記述しなければならない。
