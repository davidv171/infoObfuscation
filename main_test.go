package main

import (
	"reflect"
	"testing"
)

func Test_reconstruct3D(t *testing.T) {
	type args struct {
		deserialized []uint32
	}
	var tests = []struct {
		name string
		args args
		want [][][]uint32
	}{
		{"Basic test", args{[]uint32{
			0, 1, 8, 16, 9, 2, 3, 10,
			17, 24, 32, 25, 18, 11, 4, 5,
			12, 19, 26, 33, 40, 48, 41, 34,
			27, 20, 13, 6, 7, 14, 21, 28,
			35, 42, 49, 56, 57, 50, 43, 36,
			29, 22, 15, 23, 30, 37, 44, 51,
			58, 59, 52, 45, 38, 31, 39, 46,
			53, 60, 0, 0, 0, 0, 0, 0,

			0, 1, 8, 16, 9, 2, 3, 10,
			17, 24, 32, 25, 18, 11, 4, 5,
			12, 19, 26, 33, 40, 48, 41, 34,
			27, 20, 13, 6, 7, 14, 21, 28,
			35, 42, 49, 56, 57, 50, 43, 36,
			29, 22, 15, 23, 30, 37, 44, 51,
			58, 59, 52, 45, 38, 31, 39, 46,
			53, 60, 0, 0, 0, 0, 0, 0,

			0, 1, 8, 16, 9, 2, 3, 10,
			17, 24, 32, 25, 18, 11, 4, 5,
			12, 19, 26, 33, 40, 48, 41, 34,
			27, 20, 13, 6, 7, 14, 21, 28,
			35, 42, 49, 56, 57, 50, 43, 36,
			29, 22, 15, 23, 30, 37, 44, 51,
			58, 59, 52, 45, 38, 31, 39, 46,
			53, 60, 0, 0, 0, 0, 0, 0,
		}}, [][][]uint32{
			{{0, 1, 8, 16, 9, 2, 3, 10},
				{17, 24, 32, 25, 18, 11, 4, 5},
				{12, 19, 26, 33, 40, 48, 41, 34},
				{27, 20, 13, 6, 7, 14, 21, 28},
				{35, 42, 49, 56, 57, 50, 43, 36},
				{29, 22, 15, 23, 30, 37, 44, 51},
				{58, 59, 52, 45, 38, 31, 39, 46},
				{53, 60, 0, 0, 0, 0, 0, 0}},

			{{0, 1, 8, 16, 9, 2, 3, 10},
				{17, 24, 32, 25, 18, 11, 4, 5},
				{12, 19, 26, 33, 40, 48, 41, 34},
				{27, 20, 13, 6, 7, 14, 21, 28},
				{35, 42, 49, 56, 57, 50, 43, 36},
				{29, 22, 15, 23, 30, 37, 44, 51},
				{58, 59, 52, 45, 38, 31, 39, 46},
				{53, 60, 0, 0, 0, 0, 0, 0}},

			{{0, 1, 8, 16, 9, 2, 3, 10},
				{17, 24, 32, 25, 18, 11, 4, 5},
				{12, 19, 26, 33, 40, 48, 41, 34},
				{27, 20, 13, 6, 7, 14, 21, 28},
				{35, 42, 49, 56, 57, 50, 43, 36},
				{29, 22, 15, 23, 30, 37, 44, 51},
				{58, 59, 52, 45, 38, 31, 39, 46},
				{53, 60, 0, 0, 0, 0, 0, 0}},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstruct3D(tt.args.deserialized); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstruct3D() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}

func Test_deserialize(t *testing.T) {
	bytesds := make([]byte, 64*4)
	c := 1
	for i := 0; i < len(bytesds); i += 4 {
		bytesds[i+3] = 0
		bytesds[i+2] = 0
		bytesds[i+1] = 0
		bytesds[i] = byte(c)
		c++
	}

	type args struct {
		bytesd []byte
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{"Things to serialize", args{bytesds}, []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserialize(tt.args.bytesd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serialize(t *testing.T) {
	type args struct {
		quantized  [][][]uint32
		dimensions BitmapDimensions
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serialize(tt.args.quantized, tt.args.dimensions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
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
		//{"Basic test", args{}, [][][]uint32{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blockarize(tt.args.quantized, tt.args.candidateblocks, tt.args.candidates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blockarize() = %v, want %v", got, tt.want)
			}
		})
	}
}
