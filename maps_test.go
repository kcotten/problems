package main

import (
	"reflect"
	"testing"
)

func Test_minUniqueSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test1", args{"hello"}, []string{"hel", "lo"}},
		{"test2", args{"aan"}, []string{"a", "an"}},
		{"test3", args{"trouble"}, []string{"trouble"}},
		{"test4", args{"naa"}, []string{"na", "a"}},
		{"test5", args{"helloo"}, []string{"hel", "lo", "o"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minUniqueSubstrings(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minUniqueSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minUniqueSubstrings2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test1", args{"hello"}, []string{"hel", "lo"}},
		{"test2", args{"aan"}, []string{"a", "an"}},
		{"test3", args{"trouble"}, []string{"trouble"}},
		{"test4", args{"naa"}, []string{"na", "a"}},
		{"test5", args{"helloo"}, []string{"hel", "lo", "o"}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minUniqueSubstrings2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minUniqueSubstrings2() = %v, want %v", got, tt.want)
			}
		})
	}
}
