package main

import (
	"reflect"
	"testing"
)

func Test_zeroOutMatrix(t *testing.T) {
	type args struct {
		m [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Test1",
			args{
				[][]int{
					{1, 1},
					{1, 0},
				},
			},
			[][]int{
				{1, 0},
				{0, 0},
			},
		},
		{"Test2",
			args{
				[][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			[][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroOutMatrix(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zeroOutMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
