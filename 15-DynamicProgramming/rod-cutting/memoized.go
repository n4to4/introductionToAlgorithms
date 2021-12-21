package main

func MemoizedCutRod(p []int, n int) int {
	r := make([]int, n+1)
	for i := 0; i <= n; i++ {
		r[i] = -1
	}
	return memoizedCutRodAux(p, n, r)
}

func memoizedCutRodAux(p []int, n int, r []int) int {
	if r[n] >= 0 {
		return r[n]
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

func BottomUpCutRod(p []int, n int) int {
	r := make([]int, n+1)
	r[0] = 0
	for j := 1; j <= n; j++ {
		q := -1
		for i := 1; i <= j; i++ {
			rr := p[i] + r[j-i]
			if rr > q {
				q = rr
			}
		}
		r[j] = q
	}
	return r[n]
}
