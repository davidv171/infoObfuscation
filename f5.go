package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
)

/*

   Izberete M naključnih unikatnih trojic neničelnih koeficientov srednjih frekvenc (indeksi od 4 do 32, v primeru večjega THR se ta razpon zmanjša!).
Trojice med sabo nimajo prekrivanja.
   Vsako trojico določajo koeficienti AC1, AC2 in AC3, pri čemer korespodenčni LSB biti danih koeficentov so definirani kot C1=LSB(AC1), C2=LSB(AC2) in C3=LSB(AC3).

   Za vsako trojico (C1,C2,C3) vzamite 2 bita binariziranega sporočila, definirana kot x1 in x2 in izvedite naslednje operacije za skrivanje x1 in x2:


*/
/*
Takes a quantized block(quantized: 64 - THR bits are 0), that hasn't been F5'd before, performs F5 algorithm on it
*/
type Index struct {
	x,y int
}
type Span struct {
	start int
	end   int
}

type F5 struct {
	x1, x2, c1, c2, c3 bool
	triplet            Triplet
}

//A randomly picked triplet of three neighboring coefficients
type Triplet struct {
	a, b, c uint32
}

/*Calculate the F5 algorithm on a block and a part of the text*/
func f5(command Command, block []uint32, textbits []bool) ([][]uint32) {
	//Izberi M naključnih unikatnih trojic neničelnih koeficientov srednjih frekvenc
	triplets, indexes := triplets(command, block)

	//Keeps track of textbits inside the byte
	textbitc := 0
	//Keeps track of textbits
	for i := 0; i < len(triplets); i += 3 {
		currentTriplet := Triplet{triplets[i], triplets[i+1], triplets[i+2]}

		//For each triplet take 2 textbits of the binarized message, defined as x1 and x2
		x1, x2 := textbits[textbitc], textbits[textbitc+1]
		//Every triplet is defined by 3 coefficients AC1,AC2,AC3
		//C1 = LSB(AC1), C2 = LSB(AC2), C3 = LSB(AC3)
		f5 := F5{x1, x2, false, false, false, currentTriplet}
		f5.tripletmath()

		//Change the block coefficients that were chosen with the bits

		block[indexes[i]] = f5.triplet.a
		block[indexes[i+1]] = f5.triplet.b
		block[indexes[i+2]] = f5.triplet.c
		textbitc += 2
	}
	newbl := reconstructuint(block)
	return newbl

}

func (f5 *F5) tripletmath() {
	f5.c1 = lsb(f5.triplet.a)
	f5.c2 = lsb(f5.triplet.b)
	f5.c3 = lsb(f5.triplet.c)
	//x1 = c1 + c2  && x2 = c2 + c3 -> no change
	//x1 != c1 + c2 && x2 = c2 +c3 -> negate LSB AC1
	//x1 = c1 + c2 && x2!= c2 + c3 -> negate lsb ac3
	//x1 != c1+c2 && x2 != c2 +c3 -> negate lsb ac2
	if f5.x1 == (f5.c1 != f5.c2) && f5.x2 == (f5.c2 != f5.c3) {
		//no changes
		return
	}
	if f5.x1 != (f5.c1 != f5.c2) && f5.x2 == (f5.c2 != f5.c3) {
		//negate lsb ac1
		f5.c1 = !f5.c1
		f5.triplet.a = toggleUintLSB(f5.triplet.a)
		return
	}
	if f5.x1 == f5.c1 != f5.c2 && f5.x2 != f5.c2 != f5.c3 {
		//negate lsb ac3
		f5.c2 = !f5.c2
		f5.triplet.c = toggleUintLSB(f5.triplet.c)
		return

	}
	if f5.x1 != f5.c1 != f5.c2 && f5.x2 != f5.c2 != f5.c3 {
		//negate lsb ac2
		f5.c2 = !f5.c2
		f5.triplet.b = toggleUintLSB(f5.triplet.b)
		return

	}

}



//Get the bits in the message, i is used as a counter so we don't have to keep state
func getbits(text byte) ([]bool) {
	bits := make([]bool, 8)
	var i uint8
	//Loop through the byte and turn it into bit sequence using AND and masking
	//Using an unsigned integer, so
	//7 -> 0
	for i = 0; i < 8; i++ {
		mask := byte(1 << i)
		if (text & mask) > 0 {
			bits[7-i] = true
		} else {
			bits[7-i] = false
		}

	}
	return bits

}
func text2bits(text []byte) []bool {
	bits := make([]bool, 0)
	for b := range text {
		bits = append(bits, getbits(text[b])...)
	}
	return bits
}

func bitSliceToByte(bitSlice []bool) byte {
	var i uint8 = 0
	var resultingByte byte

	for i = 0; i < uint8(len(bitSlice)); i++ {
		if (bitSlice)[i] {
			resultingByte |= 1 << (7 - i)
		}
	}
	return resultingByte
}


//Picks tripletsnum from a block
func triplets(command Command, block []uint32) ([]uint32, []int) {
	rand.Seed(int64(command.seed))
	//4-32, bigger the thr, smaller the span
	span := Span{4, 32}
	if command.thr > 32 {
		span.end = 64 - int(command.thr)
	}
	//Keep track of what we pickedindexes
	pickedindexes := make([]int, 0)
	//Pick an M amount of tripletsnum
	for i := 0; i < int(command.tripletsnum)*3; i += 3 {
		random := rng(span)
		for j := 0; j < len(pickedindexes); j++ {
			if pickedindexes[j] == random || pickedindexes[j] == random+1 || pickedindexes[j] == random+2 {
				random = rng(span)
				j = 0
			}
		}
		pickedindexes = append(pickedindexes, random)
		pickedindexes = append(pickedindexes, random+1)
		pickedindexes = append(pickedindexes, random+2)

	}
	pickedtriplets := make([]uint32, len(pickedindexes))
	for z := range pickedindexes {
		pickedtriplets[z] = block[pickedindexes[z]]
	}

	return pickedtriplets, pickedindexes

}

//generate random int between min and max val
func rng(span Span) int {

	return rand.Intn(span.end-span.start+1) + span.start
}


//Get least significant bit of a float32
func lsb(f uint32) bool {
	if f%2 == 0 {
		return false
	} else {
		return true
	}
}

//Set an LSB of a byte
func setlsb(bit bool, source byte) (byte) {
	if bit {
		return source | 1
	} else {
		return source &^ (1)
	}

}
func togglelsb(source byte) (byte) {
	source ^= 1
	return source
}

//Take in a float, toggle it's LSB and return the new value
func toggleUintLSB(u uint32) uint32 {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, u)
	if err != nil {
		fmt.Println("binary.Write failed when toggling float LSB:", err)
	}
	bytes := buf.Bytes()
	bytes[3] = togglelsb(buf.Bytes()[3])
	bits := binary.BigEndian.Uint32(bytes)
	return bits
}