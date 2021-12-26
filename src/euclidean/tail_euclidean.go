package euclidean

func TailEuclidean(n, m uint64) uint64 {
	if m == 0 {
		return n
	}

	return Euclidean(m, n%m)
}
