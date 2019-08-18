package main

import (
	"bufio"
	"fmt"
	"golang.org/x/image/bmp"
	"image"
	"os"
)

//Open the bitmap in input read
func bitmapr(path string) ([][]float32, BitmapDimensions) {
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
	x := btmp.Bounds().Size().X
	y := btmp.Bounds().Size().Y
	f.Close()
	return bmpfr(btmp), BitmapDimensions{x, y}
}

/*
bmpfr reads the pixels from the decoded bitmap into a 2D array
*/
func bmpfr(btmp image.Image) ([][]float32) {

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
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file")
	}
	size, err := f.Write(bytes)
	if err != nil {
		fmt.Println("Error writing bytes")
	}
	defer f.Sync()

	fmt.Println("Written", size, "bytes")

}

func filer(path string) ([]byte) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Cannot find binary", path)
	}
	stat, err := f.Stat()
	if err != nil {
		fmt.Println("Couldn't get file size")
	}
	r := bufio.NewReader(f)
	p := make([]byte, stat.Size())
	s, err := r.Read(p)
	if err != nil {
		fmt.Println("Could not read file")
	}
	fmt.Println("Input size", s)
	f.Close()
	return p
}
