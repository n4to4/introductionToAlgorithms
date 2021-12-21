package main

var prices []int = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30 /*10*/, 35, 40, 45, 50, 55}

func CutRod(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := -1
	for i := 1; i <= n; i++ {
		if r := p[i] + CutRod(p, n-i); r > q {
			q = r
		}
	}
	return q
}

func main() {
	_ = prices
}
