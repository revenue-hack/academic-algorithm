package main

import (
	"fmt"

	"github.com/revenue-hack/academic-algorithm/src/euclidean"
)

func main() {
	n := uint64(10)
	m := uint64(20)

	fmt.Printf("normal: m=%d, n=%d, ans=%d\n", n, m, euclidean.Euclidean(n, m))
	fmt.Printf("tail: m=%d, n=%d, ans=%d\n", n, m, euclidean.TailEuclidean(n, m))
}
