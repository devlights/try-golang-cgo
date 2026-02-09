package main

/*
#include <stdio.h>

void hello(void) {
	printf("[C] Hello World\n");
}

int add(int x, int y) {
	return x + y;
}
*/
import "C"
import (
	"fmt"
	"runtime/metrics"
	"sync"
)

type (
	_Metrics struct {
		samples []metrics.Sample
	}
)

func NewMetrics() *_Metrics {
	m := _Metrics{samples: []metrics.Sample{{Name: "/cgo/go-to-c-calls:calls"}}}
	return &m
}

func (me *_Metrics) Print(label string) {
	// 現在のメトリクスを読み込み
	metrics.Read(me.samples)

	fmt.Printf("\n=== %s ===\n", label)
	for _, sample := range me.samples {
		//
		// 一つしか指定していないが、一応チェック
		//
		if sample.Value.Kind() == metrics.KindUint64 {
			fmt.Printf("%s: %d\n", sample.Name, sample.Value.Uint64())
		} else {
			fmt.Printf("%s: (利用不可)\n", sample.Name)
		}
	}
}

func main() {
	const (
		CallCount1 = 5
		CallCount2 = 10
	)
	var (
		m     = NewMetrics()
		c     = make(chan C.int)
		drain = func() {
			for v := range c {
				fmt.Printf("[Go] z = %d\n", int(v))
			}
			fmt.Printf("[Go] DONE\n")
		}

		wg sync.WaitGroup
	)
	wg.Go(drain)

	//
	// 以下の最初のメトリクス取得時の結果は 0 ではなく 1 となる。
	//
	// これは恐らくGoのランタイムが import "C" としたタイミングでC言語側との相互運用するための
	// セットアップのようなものが行われていると推測。(_cgo_initとか)
	//
	// つまり、/cgo/go-to-c-calls:calls というメトリクスは、ユーザーが明示的に書いた C.xxx() 呼び出しだけでなくて
	// Goランタイム内部が行うC呼び出しも含めた「GoからCへの遷移の総数」をカウントしているものということになる。
	//
	m.Print("01.Init")
	{
		// C側のhello関数をN回呼び出し
		for range CallCount1 {
			C.hello()
		}
	}
	m.Print("02.After C.hello()")
	{
		// C側のadd関数をN回呼び出し
		for i := range CallCount2 {
			var (
				x = C.int(i)
				y = C.int(i + 1)
				z C.int
			)
			z = C.add(x, y)
			c <- z
		}
	}
	m.Print("03.After C.add()")

	close(c)
	wg.Wait()
}

/*
$ task
task: [default] go run main.go

=== 01.Init ===
/cgo/go-to-c-calls:calls: 1
[C] Hello World
[C] Hello World
[C] Hello World
[C] Hello World
[C] Hello World

=== 02.After C.hello() ===
/cgo/go-to-c-calls:calls: 6
[Go] z = 1
[Go] z = 3
[Go] z = 5
[Go] z = 7
[Go] z = 9
[Go] z = 11
[Go] z = 13
[Go] z = 15
[Go] z = 17
[Go] z = 19

=== 03.After C.add() ===
/cgo/go-to-c-calls:calls: 16
[Go] DONE
*/
