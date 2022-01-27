package sort_test

import (
	"testing"

	"github.com/revenue-hack/academic-algorithm/src/sort"
)

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
