package main

import (
	"reflect"
	"testing"
)

func Test_inversef5(t *testing.T) {
	type args struct {
		coeffblock []uint32
		command    Command
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"Basic testt", args{[]uint32{0,
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
			0},Command{"b","b","b",30,3,90}}, []bool{true,true,true,true,true,true}},
		{"Zeroed test", args{[]uint32{0,
			0, 0,
			00, 0, 0,
			0, 00, 00, 00,
			00, 00, 00, 00, 0,
			0, 00, 00, 00, 00, 00,
			00, 00, 00, 00, 00, 00, 0,
			0, 00, 00, 00, 00, 00, 00, 00,
			00, 00, 00, 00, 00, 00, 00,
			00, 00, 00, 00, 00, 00,
			00, 00, 00, 00, 00,
			00, 00, 00, 00,
			0, 0, 0,
			0, 0,
			0,
		},Command{"b","b","b",30,3,90}}, []bool{false,false,false,false,false,false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inversef5(tt.args.coeffblock, tt.args.command); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inversef5() = %v, want %v", got, tt.want)
			}
		})
	}
}
