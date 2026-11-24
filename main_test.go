package main

import "testing"

func TestLastStoneWeight(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1}, 1},
		{[]int{3, 3}, 0}, // full annihilation
		{[]int{}, 0},
		{[]int{10, 4}, 6},
		{[]int{1, 1, 1}, 1},
	}
	for _, c := range cases {
		if got := LastStoneWeight(c.in); got != c.want {
			t.Errorf("LastStoneWeight(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
