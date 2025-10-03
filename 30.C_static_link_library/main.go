package main

/*
#cgo CFLAGS: -g3 -O0 -Wall -Wextra -fPIC
#cgo CPPFLAGS: -I.
*/
import "C"

//export cgo_func
func cgo_func(x, y C.int, z *C.int) C.int {
	var (
		goX = int(x)
		goY = int(y)
		ans = goX + goY
	)

	*z = C.int(ans)

	return C.int(0)
}

func main() {}
