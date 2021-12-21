package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	name string
	n    int
	want int
}

var tests []TestCase = []TestCase{
	{
		n:    1,
		want: 1,
	},
	{
		n:    2,
		want: 5,
	},
	{
		n:    3,
		want: 8,
	},
	{
		n:    5,
		want: 13,
	},
	{
		n:    6,
		want: 17,
	},
	{
		n:    7,
		want: 18,
	},
	{
		n:    10,
		want: 30,
	},
}

func TestMemoizedCutRod(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MemoizedCutRod(prices, tt.n); got != tt.want {
				t.Errorf("MemoizedCutRod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBottomUpCutRod(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BottomUpCutRod(prices, tt.n); got != tt.want {
				t.Errorf("BottomUpCutRod() = %v, want %v", got, tt.want)
			}
		})
	}
}

var input []int = []int{5, 10, 15}

func BenchmarkCutRod(b *testing.B) {
	for _, in := range input {
		b.Run(fmt.Sprintf("CutRod %d", in), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BottomUpCutRod(prices, in)
				CutRod(prices, in)
			}
		})
	}
}

func BenchmarkMemoizedCutRod(b *testing.B) {
	for _, in := range input {
		b.Run(fmt.Sprintf("Memoized %d", in), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MemoizedCutRod(prices, in)
			}
		})
	}
}

func BenchmarkBottomUpCutRod(b *testing.B) {
	for _, in := range input {
		b.Run(fmt.Sprintf("BottomUp %d", in), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BottomUpCutRod(prices, in)
			}
		})
	}
}
