# try-golang-cgo

[try-golang](https://github.com/devlights/try-golang) プロジェクトの姉妹版。cgoに関連しているサンプルが配置されています。

## サンプルリスト

| サンプル | 概要 |
| --- | --- |
| [01.CGO_Header](./01.CGO_Header/) | Cのコードをコメントヘッダに記述する基本的な方法 |
| [02.Restrictions](./02.Restrictions/) | cgoの制約事項（サポートされていないCの機能など）について |
| [03.C_int](./03.C_int/) | Cの `int` 型とGoの `int` 型を相互に変換する方法 |
| [04.C_struct](./04.C_struct/) | Cの `struct` をGoで利用する方法 |
| [05.C_string](./05.C_string/) | Goの文字列をCの文字列 (`C.CString`) に変換する方法 |
| [06.C_GoString](./06.C_GoString/) | Cの文字列をGoの文字列 (`C.GoString`) に変換する方法 |
| [07.C_CBytes](./07.C_CBytes/) | GoのバイトスライスをCの `void*` (`C.CBytes`) として渡す方法 |
| [08.C_GoBytes](./08.C_GoBytes/) | Cのデータと長さをGoのバイトスライス (`C.GoBytes`) に変換する方法 |
| [09.C_GoStringN](./09.C_GoStringN/) | 長さを指定してCの文字列をGoの文字列 (`C.GoStringN`) に変換する方法 |
| [10.C_ByteSliceToVoidPtr](./10.C_ByteSliceToVoidPtr/) | Goのバイトスライスを `void*` としてCに渡す方法 |
| [11.C_ByteSliceToCharPtr](./11.C_ByteSliceToCharPtr/) | Goのバイトスライスを `char*` としてCに渡す方法 |
| [12.C_malloc](./12.C_malloc/) | Cの `malloc` でメモリを確保し `free` で解放する方法 |
| [13.C_PointerArithmetic](./13.C_PointerArithmetic/) | Cのポインタ演算をGoで行う方法 |
| [14.C_ExportGoFunc](./14.C_ExportGoFunc/) | Goの関数をエクスポート (`//export`) してCから呼び出す方法 |
| [15.C_CallGoFunctionFromC_via_FunctionPointer](./15.C_CallGoFunctionFromC_via_FunctionPointer/) | CからGoの関数を関数ポインタ経由で呼び出す方法 |
| [16.C_dlopen_dlsym](./16.C_dlopen_dlsym/) | `dlopen` と `dlsym` を使って動的に共有ライブラリを扱う方法 |
| [17.C_init_func](./17.C_init_func/) | Cのコンストラクタ関数 (`__attribute__((constructor))`) の動作確認 |
| [18.C_init_func_so](./18.C_init_func_so/) | 共有ライブラリにおけるCのコンストラクタ関数の動作確認 |
| [19.C_NULL](./19.C_NULL/) | Cの `NULL` ポインタの扱い方 |
| [20.C_CFLAGS_LDFLAGS](./20.C_CFLAGS_LDFLAGS/) | `CFLAGS` と `LDFLAGS` を使って外部ライブラリをリンクする方法 |
| [21.C_CC](./21.C_CC/) | 使用するCコンパイラを `CC` で指定する方法 |
| [22.C_SRCDIR](./22.C_SRCDIR/) | `SRCDIR` を使って別ディレクトリのCソースファイルをビルド対象に含める方法 |
| [23.C_FixedSizeCharArrayToPointer](./23.C_FixedSizeCharArrayToPointer/) | Cの固定長文字配列をポインタとして扱う方法 |
| [24.C_Using_cgo_Handle](./24.C_Using_cgo_Handle/) | `cgo.Handle` を使ってGoの値を安全にCとやり取りする方法 |
| [25.C_struct_not_typedef](./25.C_struct_not_typedef/) | `typedef` されていない `struct` (`struct タグ名`) の扱い方 |
| [26.C_sizeof_struct](./26.C_sizeof_struct/) | Cの `struct` のサイズをGoで取得する方法 |
| [27.C_string_list](./27.C_string_list/) | CとGoの間で文字列の配列（`char**`）をやり取りする方法 |
| [28.C_Call_Go_sharedlib](./28.C_Call_Go_sharedlib/) | Goで共有ライブラリ (`.so`) を作成しCから呼び出す方法 |
| [29.C_memcpy](./29.C_memcpy/) | Cの `memcpy` を利用してメモリをコピーする方法 |
| [30.C_static_link_library](./30.C_static_link_library/) | 静的ライブラリ (`.a`) をリンクする方法 |
| [31.C_callback](./31.C_callback/) | Goの関数をCにコールバックとして渡す（基本的なイディオム） |
| [32.C_callback_safe_way](./32.C_callback_safe_way/) | Goの関数をCにコールバックとして渡す（Cの`typedef`を利用する方法） |




## 参考情報

- [C? Go? Cgo!](https://go.dev/blog/cgo)
- [Go Wiki: cgo](https://go.dev/wiki/cgo)
- [cmd/cgo](https://pkg.go.dev/cmd/cgo)
- [runtime/cgo](https://pkg.go.dev/runtime/cgo)
- [cgoを使ったCとGoのリンクの裏側 (1)](https://qiita.com/yugui/items/e71d3d0b3d654a110188)
- [cgoを使ったCとGoのリンクの裏側 (2)](https://qiita.com/yugui/items/cc490d080e0297251090)
- [ebitengine/purego](https://github.com/ebitengine/purego)
- [JupiterRider/ffi](https://github.com/JupiterRider/ffi)
