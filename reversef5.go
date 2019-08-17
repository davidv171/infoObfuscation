package main

import "fmt"

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
		btou(x1)
		btou(x2)
		//Change the block coefficients that were chosen with the bits
	}
	fmt.Println("")
	return messagebits
}
func (f5 *F5) inversetripletmath() {
	f5.c1 = lsb(f5.triplet.a)
	f5.c2 = lsb(f5.triplet.b)
	f5.c3 = lsb(f5.triplet.c)
	f5.x1 = f5.c1 != f5.c2
	f5.x2 = f5.c2 != f5.c3
}

func btou(b bool) {
	if b {
		fmt.Print(1, " ")
	} else {
		fmt.Print(0, " ")

	}
}
