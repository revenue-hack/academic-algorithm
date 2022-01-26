package sort

func BubbleSort(count int, numbers []int) []int {
	for i := 0; i < count; i++ {
		if i == count-1 {
			break
		}

		for j := 0; j < count-i; j++ {
			if cmp, cmp2 := numbers[j], numbers[j+1]; cmp > cmp2 {
				numbers[j+1] = cmp
				numbers[j] = cmp2
			}
			if j == count-2 {
				break
			}
		}
	}

	return numbers
}
