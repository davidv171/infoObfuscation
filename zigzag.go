package main

//Global zig zag, takes in an array of 8x8 blocks, then zig zags all of them simultaneously, offsetting the indexes by each matrix index
//arrayOfBlocks[blockIndex][i][j] -> result[blockIndex + iterator]
//Quantize last THR coefficients to 0


/*Zig zag an 8x8 block*/
func blockzigzag(block [][]float32) []float32 {
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
	return result

}


/*
Last threshold amount of elements are set to 0
Threshold is set by the user
*/
func quantize(thr uint64, zigzagged []float32) ([]uint32) {
	quantized := make([]uint32,len(zigzagged))
	for i := 0; i < len(zigzagged) - int(thr); i++ {
		quantized[i] = uint32(zigzagged[i])
	}
	for j := 63; uint64(j) > uint64(63)-thr; j-- {
		quantized[j] = 0
	}
	return quantized

}
