package euclidean_test

import (
	"math/rand"
	"testing"

	"github.com/revenue-hack/academic-algorithm/src/euclidean"
)

func Benchmark_Normal(b *testing.B) {
	r := rand.Uint64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		euclidean.Euclidean(uint64(i), r)
	}
}
func Benchmark_Tail(b *testing.B) {
	r := rand.Uint64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		euclidean.TailEuclidean(uint64(i), r)
	}
}
