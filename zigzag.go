package main

import (
	"fmt"
	"os"
)

//Global zig zag, takes in an array of 8x8 blocks, then zig zags all of them simultaneously, offsetting the indexes by each matrix index
//arrayOfBlocks[blockIndex][i][j] -> result[blockIndex + iterator]
//Quantize last THR coefficients to 0

/*Zig zag an 8x8 block*/
func blockzigzag(block [][]float32) [][]float32 {
	result := make([]float32, 64)

	iterator := 0
	firsthalf := true
	i := 0
	j := 0
	for iterator < 64 {
		if i == 8 && j == 0 {
			i = 7
			j++
			firsthalf = !firsthalf
		} else if i == 8 && j == 7 {
			i = 7
		}
		if firsthalf {
			if j%2 != 0 || i%2 != 0 {
				for j != 0 {
					result[iterator] = block[i][j]
					iterator++
					i++
					j--
				}
				result[iterator] = block[i][j]
				iterator++
				i++
			} else if i%2 == 0 {
				for i != 0 {
					result[iterator] = block[i][j]
					iterator++
					i--
					j++

				}
				result[iterator] = block[i][j]
				iterator++
				j++
			}
			//Second half of the matrix
		} else {
			if j%2 != 0 && i == 7 {
				//Go up the matrix until you reach the inverse(from 5,0 -> 0,5)
				for j != 7 {
					result[iterator] = block[i][j]
					iterator++
					i--
					j++

				}
				result[iterator] = block[i][j]
				iterator++
				i++
			} else if i%2 == 0 {
				for i != 7 {
					result[iterator] = block[i][j]
					iterator++
					j--
					i++
				}
				result[iterator] = block[i][j]
				iterator++
				j++
			}
		}
	}
	//Reconstruct 2D block
	return reconstructfloat(result)

}

/*
Last threshold amount of elements are set to 0
Threshold is set by the user
*/
func quantize(thr uint64, zigzagged [][]float32) ([][]uint32) {
	if zigzagged == nil {
		fmt.Println("Did not receive valid zigzag")
		os.Exit(1)
	}
	quantized := make([][]uint32, 8)
	for i := range quantized {
		quantized[i] = make([]uint32, 8)
	}
	start := 64 - int(thr)
	index := 0
	for i := range zigzagged {
		for j := range zigzagged[0] {
			if index > start {
				quantized[i][j] = 0
			} else {
				quantized[i][j] = uint32(zigzagged[i][j])
				index++

			}
		}
	}
	return quantized

}

//Reconstruct an 8x8 matrix out of 64 length 1D matrix
func reconstructfloat(result []float32) [][]float32 {
	reconstructed := make([][]float32, 8)
	for i := range reconstructed {
		reconstructed[i] = make([]float32, 8)
	}
	c := 0
	for i := 0; i < len(result); i += 8 {
		adding := result[i : i+8]
		reconstructed[c] = adding
		c++
	}
	return reconstructed
}

//Reconstruct an 8x8 matrix out of 64 length 1D matrix
func reconstructuint(result []uint32) [][]uint32 {
	reconstructed := make([][]uint32, 8)
	for i := range reconstructed {
		reconstructed[i] = make([]uint32, 8)
	}
	c := 0
	for i := 0; i < len(result); i += 8 {
		adding := result[i : i+8]
		reconstructed[c] = adding
		c++
	}
	return reconstructed
}

//turn 8x8 into 64 length 1D
func flatten(input [][]uint32) []uint32 {
	flattened := make([]uint32, 64)
	counter := 0
	for i := range input {
		for j := range input[0] {
			flattened[counter] = input[i][j]
			counter++
		}
	}
	return flattened
}

func globalZigZag(block [][][]float32) []float32 {
	//Resulting a NxN length 1D matrix from an 8x8 block
	l := len(block)
	result := make([]float32, l*64)
	for z := 0; z < l; z++ {
		i := 1
		j := 0
		iterator := 2 * l
		result[z] = block[z][0][0]
		result[z+l] = block[z][0][1]
		//Last one is predetermined
		result[(63*l)+z] = block[z][7][7]
		firsthalf := true
		for iterator+z < 64*l {
			if i == 8 && j == 0 {
				i = 7
				j++
				firsthalf = !firsthalf
			} else if i == 8 && j == 7 {
				i = 7
			}
			if firsthalf {
				if j%2 != 0 || i%2 != 0 {
					for j != 0 {
						result[iterator+z] = block[z][i][j]
						iterator += l
						i++
						j--
					}
					result[iterator+z] = block[z][i][j]
					iterator += l
					i++
				} else if i%2 == 0 {
					for i != 0 {
						result[iterator+z] = block[z][i][j]
						iterator += l
						i--
						j++

					}
					result[iterator+z] = block[z][i][j]
					iterator += l
					j++
				}
				//Second half of the matrix
			} else {
				if j%2 != 0 && i == 7 {
					//Go up the matrix until you reach the inverse(from 5,0 -> 0,5)
					for j != 7 {
						result[iterator+z] = block[z][i][j]
						iterator += l
						i--
						j++

					}
					result[iterator+z] = block[z][i][j]
					iterator += l
					i++
				} else if i%2 == 0 {
					for i != 7 {
						result[iterator+z] = block[z][i][j]
						iterator += l
						j--
						i++
					}
					result[iterator+z] = block[z][i][j]
					iterator += l
					j++
				}
			}
		}
	}
	return result
}
