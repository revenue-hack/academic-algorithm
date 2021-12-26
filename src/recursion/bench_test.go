package recursion_test

import (
	"math/rand"
	"testing"

	"github.com/revenue-hack/academic-algorithm/src/recursion"
)

/*
cpu: Intel(R) Core(TM) i7-7567U CPU @ 3.50GHz
Benchmark_NormalFactorial-4     10957178       121.9 ns/op       0 B/op       0 allocs/op
Benchmark_TailFactorial-4       12408282        93.33 ns/op       0 B/op       0 allocs/op
PASS
ok      github.com/revenue-hack/academic-algorithm/src/recursion2.711s
*/

func Benchmark_NormalFactorial(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		recursion.NormalFactorial(uint64(rand.Int63n(50)))
	}
}

func Benchmark_TailFactorial(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		recursion.TailFactorial(uint64(rand.Int63n(50)), 1)
	}
}
