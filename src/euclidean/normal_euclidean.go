package euclidean

func Euclidean(n, m uint64) uint64 {
	if m == 0 {
		return n
	}

	v := Euclidean(m, n%m)
	return v
}
