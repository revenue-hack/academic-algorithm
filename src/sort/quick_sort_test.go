package sort_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/revenue-hack/academic-algorithm/src/sort"
)

const (
	COUNT = 1000
	MAX   = 1000
)

var (
	numbers []int
)

func init() {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, COUNT)
	for i := 0; i < COUNT; i++ {
		nums[i] = rand.Intn(MAX)
	}
	numbers = nums
}

func TestQuick(t *testing.T) {
	result := sort.QuickSort(numbers)

	for i, r := range result {
		if i == len(result)-1 {
			break
		}
		if r > result[i+1] {
			t.Errorf("result: %v", result)
		}
	}
}
