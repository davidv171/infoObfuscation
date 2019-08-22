package main

import (
	"encoding/binary"
	"fmt"
)

/*Receive 2D array of coefficients, perform inverse f5 on it, return reconstructed bits of the message*/
func inversef5(coeffblock []uint32, command Command) ([]bool) {
	if coeffblock == nil {
		fmt.Println("Invalid coefficient block")
	}
	messagebits := make([]bool, 0)

	triplets, _ := triplets(command, coeffblock)

	for i := 0; i < len(triplets); i += 3 {
		currentTriplet := Triplet{triplets[i], triplets[i+1], triplets[i+2]}
		//For each triplet take 2 textbits of the binarized message, defined as x1 and x2
		//Every triplet is defined by 3 coefficients AC1,AC2,AC3
		//C1 = LSB(AC1), C2 = LSB(AC2), C3 = LSB(AC3)

		f5 := F5{false, false, false, false, false, currentTriplet}
		f5.inversetripletmath()
		x1, x2 := f5.x1, f5.x2
		messagebits = append(messagebits, x1, x2)
		//Change the block coefficients that were chosen with the bits
	}
	return messagebits
}
func (f5 *F5) inversetripletmath() {
	f5.c1 = lsb(f5.triplet.a)
	f5.c2 = lsb(f5.triplet.b)
	f5.c3 = lsb(f5.triplet.c)
	f5.x1 = (f5.c1 != f5.c2)
	f5.x2 = (f5.c2 != f5.c3)
}

func btou(b bool) {
	if b {
		fmt.Print(1, " ")
	} else {
		fmt.Print(0, " ")

	}
}
func inversionF5Caller(command Command, deserialized []uint32, reconstructed [][][]uint32) []bool {
	candidates := make([]int, 0)
	//candidateblocks := make([][][]uint32, 0)
	length := make([]bool,0)
	message := make([]bool, 0)

	size := (len(deserialized) / 64)

	//Get message size
	for i := 0; i < 32/int((command.tripletsnum * 2)); i++ {
		x := generaterng(candidates,size)
		candidates = append(candidates, x)
		rec := reconstructed[x]
		coeffs := flatten(rec)
		length = append(length, inversef5(coeffs, command)...)

	}

	messageLength := uint32FromBits(length)


	for i := 0; i < int(messageLength)*8/int((command.tripletsnum * 2)); i++ {
		x := generaterng(candidates,size)
		candidates = append(candidates, x)
		rec := reconstructed[x]
		coeffs := flatten(rec)
		message = append(message, inversef5(coeffs, command)...)

	}
	return message
}
func uint32FromBits(message []bool) uint32 {
	extracted := make([]byte, 0)
	for i := 0; i < len(message); i += 8 {
		extracted = append(extracted, bitSliceToByte(message[i:i+8]))
	}
	return uint32FromBytes(extracted)
}

func uint32FromBytes(bytes []byte) uint32 {
	return binary.LittleEndian.Uint32(bytes)
}