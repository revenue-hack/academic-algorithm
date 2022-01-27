package sort_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/revenue-hack/academic-algorithm/src/sort"
)

func genTestData(count, max int) []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, COUNT)
	for i := 0; i < COUNT; i++ {
		nums[i] = rand.Intn(MAX)
	}
	return nums
}

func Benchmark_Quick(b *testing.B) {
	b.ResetTimer()
	nums := genTestData(10000, b.N)
	for i := 0; i < b.N; i++ {
		sort.QuickSort(nums)
	}
}

func Benchmark_Bubble(b *testing.B) {
	b.ResetTimer()
	c := 10000
	nums := genTestData(c, b.N)
	for i := 0; i < b.N; i++ {
		sort.BubbleSort(c, nums)
	}
}
