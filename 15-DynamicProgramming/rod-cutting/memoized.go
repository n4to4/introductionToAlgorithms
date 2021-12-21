package main

func MemoizedCutRod(p []int, n int) int {
	r := make(map[int]int)
	return memoizedCutRodAux(p, n, r)
}

func memoizedCutRodAux(p []int, n int, r map[int]int) int {
	if rr, ok := r[n]; ok {
		return rr
	}
	if n == 0 {
		return 0
	}
	q := -1
	for i := 1; i <= n; i++ {
		rr := p[i] + memoizedCutRodAux(p, n-i, r)
		if rr > q {
			q = rr
		}
	}
	return q
}
