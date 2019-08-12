package main

/*
   Turn bitmap into 8x8 blocks
*/
func eightxeight(pixels [][]float32) []float32 {
	//Count how many 8x8 blocks there can be in an x times y matrix
	x := len(pixels)
	y := len(pixels[0])
	count := (x * y) / 64
	blocks := make([][][]float32, count)
	for i := range blocks {
		blocks[i] = make([][]float32, 8)
		for j := range blocks[i] {
			blocks[i][j] = make([]float32, 8)
		}
	}
	z := 0
	//Extra iteration counter, so we can keep resetting m
	m := 0
	//TODO: Stop when reaching the end
	for d := 0; d < len(blocks); d++ {
		//A single 8x8 block
		for i := z * 8; i/8 < z+1; i++ {
			//f and g are in-block trackers
			f := 0
			for j := m * 8; j/8 < m+1; j++ {
				g := 0
				currPixel := pixels[i][j]
				blocks[d][f][g] = currPixel
				g++
			}
			f++
		}
		blocks[d] = blocksT(blocks[d])
		m++
		//Reached right corner
		//Restart algorithm, one row down
		if x/8 == m {
			z++
			m = 0

		}
		//The transformed 8x8 block
		//transformed := make([][]float32, x, y)
		//Transform 8x8 block by transforming all rows, then transforming all columns
		//get 8xY sized block
		//Get [0,0], [8,8],[8,16] etc. every 8th tile in the 2D array

		//build a matrix of 8x8 blocks
		//Transform H into orthogonal matrix-> Inverse is faster
		//Normalize each colmn of the starting matrix to length 1
	}
	return nil
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
