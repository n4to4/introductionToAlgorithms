package main

import (
	"fmt"
	"testing"
)

func TestMatrixChainOrder(t *testing.T) {
	tests := []struct {
		name string
		p    []int
		want int
	}{
		{
			p:    []int{30, 35, 15},
			want: 30 * 35 * 15,
		},
		{
			p:    dimention,
			want: 15125,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("p=%#v, want=%d", tt.p, tt.want), func(t *testing.T) {
			if got := MatrixChainOrder(tt.p); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
