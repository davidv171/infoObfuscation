package main

//Global zig zag, takes in an array of 8x8 blocks, then zig zags all of them simultaneously, offsetting the indexes by each matrix index
//arrayOfBlocks[blockIndex][i][j] -> result[blockIndex + iterator]
//Quantize last THR coefficients to 0
func zigZag(block [][][]float32) []float32 {
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

func quantizeZigZag(thr int64, zigzagged []float32) {

}
