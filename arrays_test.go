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

func Test_floodFill(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"Test1",
			args{
				[][]int{
					{1, 1},
					{1, 0},
				},
				1,
				0,
				2,
			},
			[][]int{
				{2, 2},
				{2, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeStacks(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threeStacks()
		})
	}
}

func Test_rotate2(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test",
			args{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate2(tt.args.matrix)
		})
	}
}

func Test_isCryptSolution(t *testing.T) {
	type args struct {
		crypt    []string
		solution [][]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test",
			args{
				[]string{
					"SEND",
					"MORE",
					"MONEY",
				},
				[][]string{
					{"O", "0"},
					{"M", "1"},
					{"Y", "2"},
					{"E", "5"},
					{"N", "6"},
					{"D", "7"},
					{"R", "8"},
					{"S", "9"},
				},
			},
			true,
		},
		{
			"test2",
			args{
				[]string{
					"A",
					"A",
					"A",
				},
				[][]string{
					{"A", "0"},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCryptSolution(tt.args.crypt, tt.args.solution); got != tt.want {
				t.Errorf("isCryptSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
