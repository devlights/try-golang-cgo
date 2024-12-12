package main

/*
 */
import "C"

//export add
func add(x, y C.int) C.int {
	var (
		p1  = int(x)
		p2  = int(y)
		ans = p1 + p2
	)
	return C.int(ans)
}

func main() {}
