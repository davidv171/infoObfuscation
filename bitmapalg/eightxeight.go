package bitmapalg

import "fmt"

//An 8x8 block
type (
	Block = [][]uint32
)
type (
	Blocks [][]Block
)

/*
   Turn bitmap into 8x8 blocks
*/
func EightxEight(pixels [][]float32) Blocks {
	//Count how many 8x8 blocks there can be in an x times y matrix
	x := len(pixels)
	y := len(pixels[0])
	fmt.Println("Pixels: ", len(pixels), " x " , len(pixels[0]))
	blocks := createBlocks(x, y)
	for i := range (blocks) {
		for j := range (blocks[i]) {
			for z := j/8; z < 8; z++ {
				for y := z/8; y < 8;y++ {
					blocks[i][j][z][y] = pixels[z][y]
					fmt.Print("[",z,",",y, "]")
				}
				fmt.Println()
			}
			fmt.Println("|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||")
		}
		fmt.Println("--------------------------------------------------------------------------------")
	}

	fmt.Println(blocks)
	return Blocks{}
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
