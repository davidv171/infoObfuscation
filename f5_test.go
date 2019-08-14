package main

import (
	"reflect"
	"testing"
)

func Test_triplets(t *testing.T) {
	type args struct {
		command Command
		block   []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{"Basic test 3 triplets picked", args{
			Command{
				"b",
				"b",
				"b",
				10,
				3,
				30,
			},
			[]float32{0,
				1, 8,
				16, 9, 2,
				3, 10, 17, 24,
				32, 25, 18, 11, 4,
				5, 12, 19, 26, 33, 40,
				48, 41, 34, 27, 20, 13, 6,
				7, 14, 21, 28, 35, 42, 49, 56,
				57, 50, 43, 36, 29, 22, 15,
				23, 30, 37, 44, 51, 58,
				59, 52, 45, 38, 31,
				39, 46, 53, 60,
				0, 0, 0,
				0, 0,
				0}},
			[]float32{48, 41, 34, 24, 32, 25, 5, 12, 19}},
		{"Basic test 4 triplets picked", args{
			Command{
				"b",
				"b",
				"b",
				10,
				4,
				40,
			},
			[]float32{0,
				1, 8,
				16, 9, 2,
				3, 10, 17, 24,
				32, 25, 18, 11, 4,
				5, 12, 19, 26, 33, 40,
				48, 41, 34, 27, 20, 13, 6,
				7, 14, 21, 28, 35, 42, 49, 56,
				57, 50, 43, 36, 29, 22, 15,
				23, 30, 37, 44, 51, 58,
				59, 52, 45, 38, 31,
				39, 46, 53, 60,
				0, 0, 0,
				0, 0,
				0}},
			[]float32{41, 34, 27, 33, 40, 48, 6, 7, 14, 25, 18, 11}},
		{"Basic test 4 triplets picked over 32 threshold", args{
			Command{
				"b",
				"b",
				"b",
				40,
				4,
				160,
			},
			[]float32{0,
				1, 8,
				16, 9, 2,
				3, 10, 17, 24,
				32, 25, 18, 11, 4,
				5, 12, 19, 26, 33, 40,
				48, 41, 34, 27, 20, 13, 6,
				7, 14, 21, 28, 35, 42, 49, 56,
				57, 50, 43, 36, 29, 22, 15,
				23, 30, 37, 44, 51, 58,
				59, 52, 45, 38, 31,
				39, 46, 53, 60,
				0, 0, 0,
				0, 0,
				0}},
			[]float32{40, 48, 41, 10, 17, 24, 11, 4, 5, 32, 25, 18}},
		{"Basic test 5 triplets picked over 32 threshold", args{
			Command{
				"b",
				"b",
				"b",
				40,
				5,
				120,
			},
			[]float32{0,
				1, 8,
				16, 9, 2,
				3, 10, 17, 24,
				32, 25, 18, 11, 4,
				5, 12, 19, 26, 33, 40,
				48, 41, 34, 27, 20, 13, 6,
				7, 14, 21, 28, 35, 42, 49, 56,
				57, 50, 43, 36, 29, 22, 15,
				23, 30, 37, 44, 51, 58,
				59, 52, 45, 38, 31,
				39, 46, 53, 60,
				0, 0, 0,
				0, 0,
				0}},
			[]float32{34, 27, 20, 17, 24, 32, 11, 4, 5, 40, 48, 41, 9, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triplets(tt.args.command, tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("triplets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lsb(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Basic test", args{53.0}, false},

		{"Basic test 0", args{53.1}, false},

		{"Basic test sanity", args{66.2}, false},

		{"Basic test for 1", args{19.999999}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lsb(tt.args.f); got != tt.want {
				t.Errorf("lsb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tripletmath(t *testing.T) {
	type args struct {
		f5 F5
	}
	tests := []struct {
		name string
		args args
		want F5
	}{
		{"Basic test for f5 triple math, all true" , args{F5{true,true,true,true,true}},F5{true,true,true,false,true}},
		{"All false for f5 triple math", args{F5{false,false,false,false,false}}, F5{false,false,false,false,false}},
		{"x1 and x2 false, others true", args{F5{false,false,true,true,true}},F5{false,false,true,true,true}},
		{"x1 and x2 true, c2 false, others true", args{F5{true,true,true,false,true}},F5{true,true,true,false,true}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tripletmath(tt.args.f5); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tripletmath() = %v, want %v", got, tt.want)
			}
		})
	}
}
