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
		{"Basic test 3 tripletsnum picked", args{
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
		{"Basic test 4 tripletsnum picked", args{
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
		{"Basic test 4 tripletsnum picked over 32 threshold", args{
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
		{"Basic test 5 tripletsnum picked over 32 threshold", args{
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
		{"Basic test 5 tripletsnum picked over 32 threshold", args{
			Command{
				"b",
				"b",
				"test",
				30,
				3,
				90,
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

		})
	}
}

func Test_toggleFloatLSB(t *testing.T) {
	type args struct {
		f uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		//01000010 01001000 00000000 00000000
		{"Basic test", args{50}, 51},
		{"Basic test", args{51}, 50},
		{"Basic test", args{255}, 254},
		{"Basic test", args{0}, 1},
		{"Basic test", args{200}, 201},
		{"Basic test", args{211}, 210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toggleUintLSB(tt.args.f); got != tt.want {
				t.Errorf("toggleUintLsb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_text2bits(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"Basic test 16,32,64", args{[]byte{16, 32, 64}}, []bool{
			false, false, false, true, false, false, false, false,
			false, false, true, false, false, false, false, false,
			false, true, false, false, false, false, false, false,
		}},
		{"Basic test LONG", args{[]byte{16, 32, 64, 16, 32, 64, 16, 32, 64, 16, 32, 64, 16, 32, 64, 16, 32, 64, 16, 32, 64, 16, 32, 64}}, []bool{
			false, false, false, true, false, false, false, false,
			false, false, true, false, false, false, false, false,
			false, true, false, false, false, false, false, false,
			false, false, false, true, false, false, false, false,
			false, false, true, false, false, false, false, false,
			false, true, false, false, false, false, false, false,
			false, false, false, true, false, false, false, false,
			false, false, true, false, false, false, false, false,
			false, true, false, false, false, false, false, false,
			false, false, false, true, false, false, false, false,
			false, false, true, false, false, false, false, false,
			false, true, false, false, false, false, false, false,
			false, false, false, true, false, false, false, false,
			false, false, true, false, false, false, false, false,
			false, true, false, false, false, false, false, false,
			false, false, false, true, false, false, false, false,
			false, false, true, false, false, false, false, false,
			false, true, false, false, false, false, false, false,
			false, false, false, true, false, false, false, false,
			false, false, true, false, false, false, false, false,
			false, true, false, false, false, false, false, false,
			false, false, false, true, false, false, false, false,
			false, false, true, false, false, false, false, false,
			false, true, false, false, false, false, false, false,
		}},
		{"Basic test test", args{[]byte{116, 101, 115, 116}}, []bool{
			false, true, true, true, false, true, false, false,
			false, true, true, false, false, true, false, true,
			false, true, true, true, false, false, true, true,
			false, true, true, true, false, true, false, false,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := text2bits(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ntext2bits() = \n%v, \nwant\n%v", got, tt.want)
			}
		})
	}
}

func TestF5_tripletmath(t *testing.T) {
	type fields struct {
		x1      bool
		x2      bool
		c1      bool
		c2      bool
		c3      bool
		triplet Triplet
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Basic test", fields{true, true, true, true, true, Triplet{40, 30, 50}}},
		{"Basic test", fields{true, true, true, true, true, Triplet{40, 30, 50}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f5 := &F5{
				x1:      tt.fields.x1,
				x2:      tt.fields.x2,
				c1:      tt.fields.c1,
				c2:      tt.fields.c2,
				c3:      tt.fields.c3,
				triplet: tt.fields.triplet,
			}
			f5.tripletmath()
		})
	}
}

func Test_f5(t *testing.T) {
	//01110100 01100101 01110011 01110100
	type args struct {
		command Command
		block   []uint32
		text    []bool
	}
	tests := []struct {
		name string
		args args
		want [][]uint32
	}{
		{"NBaso", args{Command{"b", "bruh", "test", 30, 3, 90},
			[]uint32{0,
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
				0},
			[]bool{false, true, true, true, false, true, false, false,
			}}, [][]uint32{{0,
			1, 8,
			16, 9, 2,
			3, 10}, {17, 24,
			33, 25, 18, 11, 4,
			5}, {12, 19, 27, 33, 40,
			48, 41, 34}, {27, 20, 13, 6,
			7, 14, 21, 28}, {35, 42, 49, 56,
			57, 50, 43, 36}, {29, 22, 15,
			23, 30, 37, 44, 51}, {58,
			59, 52, 45, 38, 31,
			39, 46}, {53, 60,
			0, 0, 0,
			0, 0,
			0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f5(tt.args.command, tt.args.block, tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("f5() = \n%v, want\n%v", got, tt.want)
			}
		})
	}
}

func Test_lsb(t *testing.T) {
	type args struct {
		f uint32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Basic test 33", args{33}, true},
		{"Basic test 25", args{25}, true},

		{"Basic test 18", args{18}, false},

		{"Basic test 44", args{44}, false},

		{"Basic test 41", args{41}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lsb(tt.args.f); got != tt.want {
				t.Errorf("lsb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setlsb(t *testing.T) {
	type args struct {
		bit    bool
		source byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setlsb(tt.args.bit, tt.args.source); got != tt.want {
				t.Errorf("setlsb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_togglelsb(t *testing.T) {
	type args struct {
		source byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := togglelsb(tt.args.source); got != tt.want {
				t.Errorf("togglelsb() = %v, want %v", got, tt.want)
			}
		})
	}
}
