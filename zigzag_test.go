package main

import (
	"reflect"
	"testing"
)

func Test_blockzigzag(t *testing.T) {
	type args struct {
		block [][]float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{"Basic 8x8 test", args{[][]float32{
			{0,1,2,3,4,5,6,7},
			{8,9,10,11,12,13,14,15},
			{16,17,18,19,20,21,22,23},
			{24,25,26,27,28,29,30,31},
			{32,33,34,35,36,37,38,39},
			{40,41,42,43,44,45,46,47},
			{48,49,50,51,52,53,54,55},
			{56,57,58,59,60,61,62,63},
		},
		},
			[]float32{0,
			1,8,
			16,9,2,
			3,10,17,24,
			32,25,18,11,4,
			5,12,19,26,33,40,
			48,41,34,27,20,13,6,
			7,14,21,28,35,42,49,56,
			57,50,43,36,29,22,15,
			23,30,37,44,51,58,
			59,52,45,38,31,
			39,46,53,60,
			61,54,47,
			55,62,
			63}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blockzigzag(tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blockzigzag() = %v,\n                          want %v", got, tt.want)
			}
		})
	}
}

func Test_quantize(t *testing.T) {
	type args struct {
		thr       uint64
		zigzagged []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{"Basic test cut 6" , args{6,
			[]float32{0,
			1,8,
			16,9,2,
			3,10,17,24,
			32,25,18,11,4,
			5,12,19,26,33,40,
			48,41,34,27,20,13,6,
			7,14,21,28,35,42,49,56,
			57,50,43,36,29,22,15,
			23,30,37,44,51,58,
			59,52,45,38,31,
			39,46,53,60,
			61,54,47,
			55,62,
			63}},[]float32{0,
			1,8,
			16,9,2,
			3,10,17,24,
			32,25,18,11,4,
			5,12,19,26,33,40,
			48,41,34,27,20,13,6,
			7,14,21,28,35,42,49,56,
			57,50,43,36,29,22,15,
			23,30,37,44,51,58,
			59,52,45,38,31,
			39,46,53,60,
			0,0,0,
			0,0,
			0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quantize(tt.args.thr, tt.args.zigzagged)
		})
	}
}
