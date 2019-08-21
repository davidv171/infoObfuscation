package main

import (
	"fmt"
	"math"
)

//File for the meta analysis needed for the report
/*Requires: inverse haar transform*/
/*
Shannon entropy calculations
PSNR calculations
Blockiness?
*/


func meta(original,new [][]uint32) {
	fmt.Println("PSNR:" , PSNR(original,new))

}

func PSNR(f,g [][]uint32) float64 {
	//PSNR = 20 * log(10) * ( MAXf / sqrt(MSE))
	//MSE = 1/mn €€||f(i,j) - g(i,j)||²
	sMSE := math.Sqrt(float64(MSE(f,g)))
	var maxf float64 = 255
	return 20 * math.Log10(maxf/sMSE)
}
func MSE(f,g [][]uint32) float32{
	//Images are identical in size
	var sub,p uint32 = 0,0

	for i := range f {
		for j := range f[i]{

		 sub = f[i][j] - g[i][j]

		 p += sub * sub

		}

	}
	m := float32(len(f))
	n := float32(len(f[0]))
	return (1/(m*n)) * float32(p)
}


func shannon (f [][]uint32) {
	//Calculate probability tables of each symbol 0-255
	fpt := probtable(f)
	//p[i][j] = probability of a pixel having gray level i
	//SUM(SUM(p[i][j]log2p[i][j]))
	var s float32
	for i := range f {
		for j := range f[i] {
			symbol := f[i][j]
			probability := fpt[symbol]
			s += probability * float32(math.Log2(float64(probability)))
		}
	}
}

func probtable (f [][]uint32) []float32{
	pt := make([]float32,255)
	for i := range f {
		for j := range f[i] {
			pt[f[i][j]]++
		}
	}
	ptu := make([]float32,0)
	for i := range pt {
		if pt[i] != 0 {
			ptu = append(ptu, pt[i])
		}
	}
	for i := range pt {
		pt[i] = pt[i]/float32(len(ptu))
	}
	return pt
}

func turnToInt(m [][]uint32) [][]int32{
	n := make([][]int32,len(m))

	for i := range n {
		n[i] = make([]int32,len(m[i]))
	}

	for i := range m {
		for j := range m[i] {

			n[i][j] = int32(int(m[i][j]))

		}
	}
	return n
}
