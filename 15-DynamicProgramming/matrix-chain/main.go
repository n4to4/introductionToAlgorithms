package main

import (
	"fmt"
	"math"
)

var dimention []int = []int{
	// 30x35, 35x15, ...
	30, 35, 15, 5, 10, 20, 25,
}

func MatrixChainOrder(p []int) int {
	n := len(p) - 1

	// m[1..n, 1..n]
	m := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		m[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == j {
				m[i][j] = 0
			} else {
				m[i][j] = math.MaxInt
			}
		}
	}

	// s[1..n-1, 2..n]
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n+1)
	}

	for l := 2; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			for k := i; k < j; k++ {
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}

	for i := 0; i <= n; i++ {
		fmt.Println(m[i])
	}

	return m[1][n]
}

func main() {}
