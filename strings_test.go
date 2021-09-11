package main

import (
	"testing"
)

func Test_unique(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{"hello"}, false},
		{"Test2", args{"helio"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique(tt.args.s); got != tt.want {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniqueNoSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{"hello"}, false},
		{"Test2", args{"helio"}, true},
		{"Test3", args{"h"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueNoSpace(tt.args.s); got != tt.want {
				t.Errorf("uniqueNoSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseCstyle(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{"hello\000"}, "olleh\000"},
		{"Test2", args{"o\000"}, "o\000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseCstyle(tt.args.s); got != tt.want {
				t.Errorf("reverseCstyle() = %v, want %v", got, tt.want)
			}
		})
	}
}
