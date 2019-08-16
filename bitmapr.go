package main

import (
	"bufio"
	"fmt"
	"golang.org/x/image/bmp"
	"image"
	"os"
)

//Open the bitmap in input read
func bitmapr(path string) [][]float32 {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Cannot bitmapr the bitmap in the path? ", path)
	}
	r := bufio.NewReader(f)
	fmt.Println("The bitmap is: ", r.Size(), "bytes long before obfuscation")
	btmp, err := bmp.Decode(r)
	if err != nil {
		fmt.Println("Could not decode bitmap ", err)
	}
	return bmpfr(btmp)
}

/*
bmpfr reads the pixels from the decoded bitmap into a 2D array
*/
func bmpfr(btmp image.Image) ([][]float32){

	x := btmp.Bounds().Size().X
	y := btmp.Bounds().Size().Y
	pixels := make([][]float32, x)
	fmt.Println("Bitmap dimensions, x: ", x, " y: ", y)
	for i := 0; i < x; i++ {
		pixels[i] = make([]float32, y)
		for j := 0; j < y; j++ {
			pix, _, _, _ := btmp.At(i, j).RGBA()
			//we're dealing with n bit depth gray pixel, the library always does 0-65635
			pix = pix >> 8
			pixels[i][j] = float32(pix)
		}
	}
	return pixels

}
//Write a bitmap
func filew(path string, bytes []byte) {
	f,err := os.Create("bitmaps/output")
	if err != nil {
		fmt.Println("Error creating file")
	}
	size,err := f.Write(bytes)
	if err != nil {
		fmt.Println("Error writing bytes")
	}
	fmt.Println("Written",size,"bytes")

}
