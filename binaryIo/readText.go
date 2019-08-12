package binaryIo

import (
	"bufio"
	"fmt"
	"os"
)

//Read text file as binary
func Textr(path string) []byte{
	fmt.Println("Trying to open text file" , path)
	file,err := os.Open(path)
	if(err != nil){
		fmt.Println("Could not open file " + path)
		os.Exit(1)
	}
	r := bufio.NewReader(file)
	bytes := make([]byte,r.Size())
	size, err := file.Read(bytes)
	if err != nil {
		fmt.Println("Could not read the file into bytes " , err)
		os.Exit(1)
	}
	fmt.Println("Read file of size", size , "into bytes before cutting" , len(bytes))
	return bytes[:size]
}
