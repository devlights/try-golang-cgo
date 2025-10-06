package main

import "C"
import (
	"fmt"
	"time"
)

//export export_func
func export_func(x, y C.int) C.int {
	fmt.Println("[Go] sleep 1sec")
	time.Sleep(1 * time.Second)
	ans := (x * y)
	fmt.Printf("[Go] x=%d, y=%d, ans=%d\n", x, y, ans)
	return ans
}
