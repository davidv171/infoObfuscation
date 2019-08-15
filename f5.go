package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"unsafe"
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

type Span struct {
	start int
	end   int
}

type F5 struct {
	x1, x2, c1, c2, c3 bool
}

func f5(command Command, block []float32, text []byte) {
	//Izberi M naključnih unikatnih trojic neničelnih koeficientov srednjih frekvenc
	triplets := triplets(command, block)
	//Keeps track of bytes, float32 -> 4 bytes per thing
	wordc := 0
	//Keeps track of bits inside the byte
	counter := 0
	//Keeps track of bits
	//bits := block2bits(block)
	for i := 0; i < len(triplets); i += 3 {
		c1 := lsb(triplets[i])
		c2 := lsb(triplets[i+1])
		c3 := lsb(triplets[i+2])
		//For each triplet take 2 bits of the binarized message, defined as x1 and x2
		x1, x2 := getbits(text[wordc])[counter], getbits(text[wordc])[counter+1]
		f5 := tripletmath(F5{x1, x2, c1, c2, c3})
		//Set LSBs of blocks

		fmt.Println(f5)
		counter += 2
		if i%8 == 0 {
			counter = 0
			wordc++
		}
		fmt.Println(c1, c2, c3)

	}
	//Every triplet is defined by 3 coefficients AC1,AC2,AC3
	//C1 = LSB(AC1), C2 = LSB(AC2), C3 = LSB(AC3)

}

func tripletmath(f5 F5) (F5) {

	//x1 = c1 + c2  && x2 = c2 + c3 -> no change
	//x1 != c1 + c2 && x2 = c2 +c3 -> negate LSB AC1
	//x1 = c1 + c2 && x2!= c2 + c3 -> negate lsb ac3
	//x1 != c1+c2 && x2 != c2 +c3 -> negate lsb ac2
	if f5.x1 != f5.c1 != f5.c2 && f5.c2 != f5.c3 {
		f5.c1 = !f5.c1
	}
	if f5.x1 == f5.c1 != f5.c2 && f5.x2 != f5.c2 != f5.c3 {
		f5.c2 = !f5.c2
	}
	if f5.x1 != f5.c1 != f5.c2 && f5.x2 != f5.c2 != f5.c3 {
		f5.c2 = !f5.c2
	}
	return f5

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

//Picks triplets from a block
func triplets(command Command, block []float32) ([]float32) {
	//4-32, bigger the thr, smaller the span
	span := Span{4, 32}
	if command.thr > 32 {
		span.end = 64 - int(command.thr)
	}
	rand.Seed(int64(command.seed))
	//Keep track of what we pickedindexes
	pickedindexes := make([]int, 0)
	//Pick an M amount of triplets
	for i := 0; i < int(command.triplets)*3; i += 3 {
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
	fmt.Println(pickedindexes)
	pickedtriplets := make([]float32, len(pickedindexes))
	for z := range pickedindexes {
		pickedtriplets[z] = block[pickedindexes[z]]
	}

	return pickedtriplets

}

//generate random int between min and max val
func rng(span Span) int {
	return rand.Intn(span.end-span.start+1) + span.start
}

//Get least significant bit of a float32
func lsb(f float32) bool {
	return (*(*[4]byte)(unsafe.Pointer(&f)))[3]<<7 == 1
}

//Set an LSB of a byte
func setlsb(bit bool, source byte) (byte) {
	if bit {
		return source | 1
	} else {
		return source &^ (1)
	}

}
func togglelsb(source byte) (byte){
	source ^= 1
	return source
}

//Take in a float, toggle it's LSB and return the new value
func toggleFloatLSB(f float32) float32 {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed when toggling float LSB:", err)
	}
	bytes := buf.Bytes()
	bytes[3] = togglelsb(buf.Bytes()[3])
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}