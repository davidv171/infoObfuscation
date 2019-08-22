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


func meta(original [][]int32,new [][]uint32) {
	intnew := turnToInt(new)
	fmt.Println("PSNR:" , PSNR(original,intnew), "dB")
	fmt.Println("Shannon's entropy of the original" , shannon(original))
	fmt.Println("Shannon's entropy of the new image" , shannon(intnew))
	fmt.Println("Blockiness of the original" , blockiness(original))
	fmt.Println("Blockiness of the new image" , blockiness(intnew))


}

func PSNR(f,g [][]int32) float64 {
	//PSNR = 20 * log(10) * ( MAXf / sqrt(MSE))
	//MSE = 1/mn €€||f(i,j) - g(i,j)||²
	sMSE := math.Sqrt(float64(MSE(f,g)))
	var maxf float64 = 255
	return 20 * math.Log10(maxf/sMSE)
}
func MSE(f,g [][]int32) float32{
	//Images are identical in size
	var sub,p int32 = 0,0

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


func shannon (f [][]int32) float32{
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
	return s
}

func probtable (f [][]int32) []float32{
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

func blockiness(g [][]int32) (float64){
	x := (len(g)-1)/8
	fmt.Print("x" , x)
	var b float64
	for i := 1; i < x; i++ {
		for j := 1; j < len(g[i]);j++ {
			d := (8*i)+1
			b+= math.Abs(float64(g[8*i][j] - g[d][j]))
		}
	}
	y := (len(g[0])-1)/8
	var c float64
	for j := 1; j < y; j++ {
		for i := 1; i < len(g);i++ {
			d := (8*j)+1
			c+= math.Abs(float64(g[i][8*j] - g[i][d]))
		}
	}
	return b+c
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


func floatToInt(m [][]float32) [][]int32{
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