package main

import (
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




