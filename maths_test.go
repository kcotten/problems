package main

import "testing"

func Test_powerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{2},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := powerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
