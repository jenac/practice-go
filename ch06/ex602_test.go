package ch06

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	type Args struct {
		v1 int
		v2 int
	}
	tests := []struct {
		name string
		args Args
		want int
	}{
		{
			name: "1+2",
			args: Args{1, 2},
			want: 3,
		},
		{
			name: "5+7",
			args: Args{5, 7},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got:= Add(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("Add() = %v, want %c", got, tt.want)
			}
		})
	}
}

func TestFloatToString(t *testing.T) {
	type Args struct {
		f1 float64
		f2 float64
	}
	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "1/3",
			args: Args{1.0, 3.0},
			want: "33.33%",
		},
		{
			name: "2/3",
			args: Args{2.0, 3.0},
			want: "66.67%",
		},
	}

	for _, tt:=range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got:=FloatToString(tt.args.f1, tt.args.f2); got != tt.want {
				t.Errorf("FloatToString: %v != %v", got, tt.want)
			}
		})
	}

}

func ExampleAdd() {
	fmt.Println(Add(100, 200))
	// Output:
	// 300
}
