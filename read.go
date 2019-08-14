package main

import (
	"fmt"
	"os"
	"strconv"
)

func Read() Command {
	if len(os.Args) != 6 {
		fmt.Println("Not enough or too many command args")
	} else {
		thr,triplets := integers()
		return Command{
			os.Args[1],
			os.Args[2],
			os.Args[3],
			thr,
			triplets,
			triplets*thr,
		}
	}
	fmt.Println("Exiting due to false input commands")
	os.Exit(1)
	return Command{}
}
func integers() (uint64,uint64){
	thr, err := strconv.ParseUint(os.Args[4], 10, 32)
	if err != nil {
		fmt.Println("Could not parse integer for threshold")
	}
	triplets, err := strconv.ParseUint(os.Args[5], 10, 32)
	if err != nil {
		fmt.Println("Could not parse integer for M")
	}
	return thr,triplets
}

/*
Input arguments:
./binary inputImage options input/output message threshold M
input file: path to input BMP file
options:
	h -> hide message
	e -> extract message
input/output message: path to input/output text message
threshold: threshold in compression
M: amount of unique triplet non zero coefficients, used in F5 steganography

*/
type Command struct {
	bmp      string
	option   string
	message  string
	thr      uint64
	triplets uint64
	seed     uint64
}
