package main

import (
	"reflect"
	"testing"
)

func Test_generateTestArray2D(t *testing.T) {
	tests := []struct {
		name string
		n, m int
	}{
		{"Basic 16 16", 16, 16},
		{"Basic 32 32", 32, 32},
		{"Basic 64 64", 64, 64},
		{"Basic 128 128", 128, 128},
		{"Basic 256 256", 256, 256},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateTestArray(tt.n, tt.m)
		})
	}
}

func Test_blockarize(t *testing.T) {
	type args struct {
		quantized       [][][]uint32
		candidateblocks [][][]uint32
		candidates      []int
	}
	tests := []struct {
		name string
		args args
		want [][][]uint32
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blockarize(tt.args.quantized, tt.args.candidateblocks, tt.args.candidates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blockarize() = %v, want %v", got, tt.want)
			}
		})
	}
}
