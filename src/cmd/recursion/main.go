package main

import (
	"fmt"

	"github.com/revenue-hack/academic-algorithm/src/recursion"
)

func main() {
	n := uint32(30)
	fmt.Printf("normal recursion: n=%d\tanswer=%d\n", n, recursion.NormalFactorial(n))

	fmt.Printf("tail recursion: n=%d\tanswer=%d\n", n, recursion.TailFactorial(n, 1))
}
