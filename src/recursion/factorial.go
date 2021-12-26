package recursion

func NormalFactorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	v := n * NormalFactorial(n-1)
	return v
}
