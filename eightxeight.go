package main

import "fmt"

/*
   Turn bitmap into 8x8 blocks
*/
func EightxEight(pixels [][]float32) [][][][]float32 {
	//Count how many 8x8 blocks there can be in an x times y matrix
	length := len(pixels)
	width := len(pixels[0])
	fmt.Println("Pixels: ", len(pixels), " x ", len(pixels[0]))
	blocks := createBlocks(length, width)
	f := 0
	s := 0
	for i := range (blocks) {
		for j := range (blocks[i]) {
			//Check for right corner
			for z := s * 8; z/8 < s+1; z++ {
				fmt.Print("|")
				for y := f * 8; y/8 < f+1; y++ {
					fmt.Print("[", z, ",", y, "]")
					b := z%8
					n := y%8
					blocks[i][j][b][n] = pixels[z][y]
				}
				fmt.Println("|")
			}
			f++
			if f == length/8 {
				f = 0
				s++
			}
			fmt.Println("f", f)
		}
		if s == width/8 {
			s = 0;
			f++
		}

		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("")

	}

	fmt.Println(blocks)
	return blocks
}

func createBlocks(x int, y int) [][][][]float32 {
	length := x / 8
	width := y / 8
	blocks := make([][][][]float32, length)
	for g := range blocks {
		blocks[g] = make([][][]float32, width)
		for i := range blocks[g] {
			blocks[g][i] = make([][]float32, 8)
			for j := range blocks[g][i] {
				blocks[g][i][j] = make([]float32, 8)
			}
		}
	}
	return blocks
}
