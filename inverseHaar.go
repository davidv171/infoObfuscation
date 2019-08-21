package main

/*Transpose 8x8 matrix*/

func inverseHaar(command Command, deserialized []uint32, dim uint32) {

	/*first uint32 is the dimension of the image, 2^n*/
	/*1. Reverse global zig zag
	2. On each block : frequency to time
		//Reverse haar on columns
	//Reverse haar on rows
	///Reverse Haar = Transpose the received block if normalized
	*/
	reversed := reverseglobalzig(deserialized)
	invhaar := make([][][]uint32, len(reversed))
	for i := range reversed {
		invhaar[i] = reversed[i]
	}
	//Rebuild the matrix ready for bitmap
	matrix := make([][]uint32, dim)
	for i := range matrix {
		matrix[i] = make([]uint32, dim)
	}
	fourd := construct4D(invhaar)
	serialized := serialize4d(fourd)

	s := turnToBytes(serialized)
	btmpfw(s, "output.bmp")

}

func turnToBytes(serialized [][]uint32) [][]byte {
	s := make([][]byte, 512)
	for i := range s {
		s[i] = make([]byte, 512)
	}
	for i := range serialized {
		for j := range serialized[i] {
			s[i][j] = byte(serialized[i][j])
		}
	}
	return s
}
func serialize4d(fourd [][][][]uint32) [][]uint32 {
	x, y := 512, 512
	d := make([][]uint32, x)
	for i := range d {
		d[i] = make([]uint32, y)
	}
	cnt := 0
	j := 0
	i := 0
	c := 0
	f := 0
	for z := 0; z < len(fourd[0]); z++ {
		//fmt.Println("[", f, ",", z+(cnt*8), "] -> [", c, "][", j, "][", i, "][", z, "]")
		d[f][z+(cnt*8)] = fourd[c][j][i][z]
		//End of the row, go to the right neighbour block, take the ith row there
		if z == 7 {
			z = -1
			j++
			cnt++

		}
		//Reached the end of the rows, restart algorithm a column down
		if j == len(fourd[c]) {
			j = 0
			z = -1
			i++
			cnt = 0
			f++

		}
		//Reached the last corner of the array of blocks, restart on the array of blocks a column down
		if i == 8 {
			z = -1
			j = 0
			i = 0
			c++

		}
		if (f == 512) {
			break
		}

	}
	return d

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
