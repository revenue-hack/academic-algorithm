package sort

func QuickSort(numbers []int) []int {
	return quick(numbers)
}

func quick(array []int) []int {
	numbers := array

	if l := len(numbers); l < 2 {
		return numbers
	}
	base := numbers[0]

	g1 := make([]int, 0, len(numbers))
	g2 := make([]int, 0, len(numbers))

	for _, num := range numbers[1:] {
		if base > num {
			g1 = append(g1, num)
		} else {
			g2 = append(g2, num)
		}
	}

	g1 = append(g1, base)
	return append(quick(g1), quick(g2)...)
}
