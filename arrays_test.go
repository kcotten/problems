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
		// TODO: Add test cases.
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threeStacks()
		})
	}
}
