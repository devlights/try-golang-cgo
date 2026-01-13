# これは何？

Go 1.23で導入された `structs.HostLayout` を使用して、C言語の構造体とGoの構造体のメモリレイアウトを完全に一致させ、安全にポインタキャスト（ゼロコピー）を行うサンプルです。

## 概要

cgoを利用する場合、C側で定義された構造体のメモリレイアウト（フィールドの並びやアライメントによるパディング）と、Go側の構造体レイアウトが一致している必要があります。

従来、Goの構造体レイアウトはコンパイラの最適化に依存しており、将来的なGoのバージョンアップにおいてフィールドの並べ替えやパディングの変更が行われないという保証はありませんでした。`structs.HostLayout` を構造体に含めることで、その構造体が**ホストプラットフォームのC ABI準拠のレイアウト**であることをコンパイラに指示できます。

本サンプルでは、あえてアライメント（境界調整）が発生する構成にしています。

## C言語側の定義

```c
typedef struct {
    uint8_t val1; // 1byte
                  // [Padding] 3bytes (int32_tのアライメントのため)
    int32_t val2; // 4byte
    int64_t val3; // 8byte
} c_struct;

```

C言語（標準的なx86_64/ARM64環境）では、`int32_t` は4バイト境界、`int64_t` は8バイト境界に配置されます。そのため、`val1`（1バイト）の直後に3バイトのパディングが自動的に挿入され、`val2` の開始オフセットが4になります。`val2`と`val3`の間は８バイト境界に収まっているのでパディングは発生しません。

## Go言語側の定義

```go
type Go_Struct struct {
    _    structs.HostLayout
    val1 uint8
    val2 int32
    val3 int64
}

```

`_ structs.HostLayout` をフィールドに配置することで、Goコンパイラに対し「この構造体は、このプラットフォームのCコンパイラと同じルールでレイアウトせよ」という制約を課しています。

## `structs.HostLayout` の有無による違い

| 比較項目 | `structs.HostLayout` あり | `structs.HostLayout` なし (標準) |
| --- | --- | --- |
| **レイアウト保証** | ホストのC ABIに準拠することを保証。 | Goコンパイラの仕様に依存（将来的に変更の可能性あり）。 |
| **パディング** | Cコンパイラと同様のパディングが挿入される。 | 効率化のためにパディングが詰められたり順序が変わる可能性がある。 |
| **安全性** | `unsafe.Pointer` によるキャストが仕様として正当化される。 | 実装上動く可能性は高いが、厳密には未定義動作のリスクを伴う。 |

## 検証内容

本プログラムを実行すると、以下の項目を出力してレイアウトの一致を確認します。

1. **値の整合性**: C側で `malloc` して値を代入したメモリを、Goのポインタとして解釈しても正しく読み取れるか。
2. **メモリアドレスとパディング**: 各フィールドのメモリアドレスの差分を計算し、`val1` と `val2` の間に意図した通りの3バイトのパディングが存在するか。

### 実行結果

```bash
$ task
task: [default] go run main.go
[C ] val1=1, val2=2, val3=3
[Go] val1=1, val2=2, val3=3
アドレス：  (1) 0x3379b050, (2) 0x3379b054, (3) 0x3379b058
パディング: (1) 3 bytes, (2) 0 bytes
```

## 参考情報

* [Go 1.23 Release Notes - Standard library](https://www.google.com/search?q=https://tip.golang.org/doc/go1.23%23structs)
* [Type structs.HostLayout - go.dev](https://www.google.com/search?q=https://pkg.go.dev/structs%23HostLayout)

