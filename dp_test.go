package main

import "testing"

func Test_knapsack(t *testing.T) {
	type args struct {
		profits  []int
		weights  []int
		capacity int
		idx      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				[]int{1, 6, 10, 16},
				[]int{1, 2, 3, 5},
				7,
				0,
			},
			22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack(tt.args.profits, tt.args.weights, tt.args.capacity, tt.args.idx); got != tt.want {
				t.Errorf("knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}
