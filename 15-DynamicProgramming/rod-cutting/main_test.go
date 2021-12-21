package main

import "testing"

func TestCutRod(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutRod(prices, tt.n); got != tt.want {
				t.Errorf("CutRod() = %v, want %v", got, tt.want)
			}
		})
	}
}
