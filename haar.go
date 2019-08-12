package main

import "math"

/* Takes an input row, calculates haar transform on it, depth elements just
get appended without calculation We send rows and columns alike as
a 1D array, and return a 1D array
which is the haar transform of it We only append, without calculating from depth onwards
First run: depth = 0-> don't append anything,
Second run: depth = 4 -> append last 4 elements without calculaitng anything
Third run depth = 6, append last 6 elements without calculating anything
number of runs is decided by log2(8) -> we're dealing with 8 sized rows/columns
*/
func haar(input []float32, thr float32, depth int) []float32 {
	//Sums and subtraction array, later we append subtract to the sums
	sums := make([]float32, 0, len(input)/2)
	subtr := make([]float32, 0, len(input)/2)
	//Input should always be just a row
	//So we're expecting 8 x of 8 sized rows to be inputted N times
	for i := 1; i < len(input)-depth; i++ {
		//Calculate averages and differences
		sqrt := float32(math.Sqrt(2))
		if !(i%2 == 0) {
			var sum = (input[i-1] + input[i]) / 2
			//Round to the 4th decimal
			sum *= sqrt
			var sub = (input[i-1] - input[i]) / 2
			sub *= sqrt
			if sum < thr {
				sum = 0
			}
			if sub < thr {
				sub = 0
			}
			//Round to 4 decimals
			sums = append(sums, float32(math.Round(float64(sum*10000))/10000))
			subtr = append(subtr, float32(math.Round(float64(sub*10000)/10000)))
		}
	}

	subtr = append(subtr, input[len(input)-depth:]...)
	rowhaar := append(sums, subtr...)
	//Unless we're at the last depth, recurse
	//Ugly as shit recursion, please FIXME
	switch depth {
	case 0:
		return haar(rowhaar, thr, 4)
	case 4:
		return haar(rowhaar, thr, 6)
	case 6:
		return rowhaar
	default:
		return nil
	}
}

func blocksT(block [][]float32, thr float32) [][]float32 {
	transformedBlock := make([][]float32, 8, 8)
	for i := 0; i < 8; i++ {
		transformedBlock[i] = haar(getRow(block, i), thr, 0)
	}
	//Transform blocks after, used the already transformed matrix...
	for i := 0; i < 8; i++ {
		//Get column as a row-> insert it as a column
		currColumn := getColumn(transformedBlock, i)
		for j := 0; j < 8; j++ {
			transformedBlock[j][i] = haar(currColumn, thr, 0)[j]
		}
	}
	//Transpose it because I have low iq
	return transformedBlock
}

//Receive 8x8 block, return 1D array of size 8, based on index value
//index 0 -> get row 0 in 8x8 block
func getRow(block [][]float32, index int) []float32 {
	row := block[index][:]
	return row
}

//Receive 8x8 block, return 1D array of size 8, representing the indexth column
func getColumn(block [][]float32, index int) []float32 {
	column := make([]float32, 8)
	//alternative := block[:][index]
	for i := range column {
		column[i] = block[i][index]
	}
	return column
}
