package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/revenue-hack/academic-algorithm/src/sort"
)

const (
	COUNT = 100
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

func main() {

	nums := numbers

	fmt.Printf("answers: %v\n", sort.BubbleSort(COUNT, nums))
	fmt.Printf("answers: %v\n", numbers)
}
