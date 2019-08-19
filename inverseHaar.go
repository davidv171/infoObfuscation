package main

import "fmt"

func inverseHaar(command Command, reconstructed [][][]uint32, dim uint32) {

	/*first uint32 is the dimension of the image, 2^n*/
	/*1. Reverse global zig zag
	2. On each block : frequency to time

	*/
	for i := range reconstructed {
		reversedz := reversezigzag(flatten(reconstructed[i]))
		fmt.Println(reversedz)

	}
}

func reversezigzag(flat []uint32) [][]uint32 {
	return nil
}
