package sort_test

import (
	"math/rand"
	"time"
)

const (
	COUNT = 10000
	MAX   = 100000
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
