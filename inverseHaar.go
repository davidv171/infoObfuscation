package main

import "fmt"

func inverseHaar(command Command, deserialized []uint32, dim uint32) {

	/*first uint32 is the dimension of the image, 2^n*/
	/*1. Reverse global zig zag
	2. On each block : frequency to time
	*/
	reversed := reverseglobalzig(deserialized)
	fmt.Println(reversed)

}

func reverseglobalzig(flat []uint32) [][][]uint32 {
	//Amount of 8x8 blocks in the entire
	delimiter := 4096
	recon := make([][][]uint32, delimiter)
	for i := range recon {
		recon[i] = make([][]uint32, 8)
		for j := range recon[i] {
			recon[i][j] = make([]uint32, 8)
		}
	}
	b := 0
	g := 0
	//d := 0
	//get them into the correct 8x8 blocks first then de-zig-zag
	for i := 0; i < len(flat); i++ {
		shift := i % (delimiter)
		if i != 0 && shift == 0 {
			b++
			if b == 8 {
				g++
				b = 0
			}
			if g == 8 {
				g = 0
			}

		}
		cand := flat[i]
		recon[shift][g][b] = cand
	}

	unzig := make([][][]uint32, len(recon))
	for i := range recon {
		unzig[i] = reverselocalzig(recon[i])
	}
	return unzig
}

//Reverse the zig zag on the block
func reverselocalzig(block [][]uint32) [][]uint32 {
	lt := lookupTable()
	zigged := make([]uint32, 64)
	flat := flatten(block)
	for i := range block {
		for j := range block[i] {
			if block[i][j] != 0 {
			}
		}
	}
	for i := 0; i < 64; i++ {
		b := lt[i]
		zigged[b] = flat[i]
	}
	return reconstructuint(zigged)
}

//Contains the lookup table with indexes, written for a bunch of test
func lookupTable() []uint32 {
	return []uint32{0, 1, 8, 16, 9, 2, 3, 10,
		17, 24, 32, 25, 18, 11, 4, 5,
		12, 19, 26, 33, 40, 48, 41, 34,
		27, 20, 13, 6, 7, 14, 21, 28,
		35, 42, 49, 56, 57, 50, 43, 36,
		29, 22, 15, 23, 30, 37, 44, 51,
		58, 59, 52, 45, 38, 31, 39, 46,
		53, 60, 61, 54, 47, 55, 62, 63}
}
