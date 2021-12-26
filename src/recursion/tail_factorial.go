package recursion

func TailFactorial(n uint64, acc uint64) uint64 {
	if n == 0 {
		return acc
	}

	return TailFactorial(n-1, n*acc)
}
