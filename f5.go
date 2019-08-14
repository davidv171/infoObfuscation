package main

import (
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

type Span struct {
	start int
	end   int
}

func f5(command Command, block []float32) {
	//Izberi M naključnih unikatnih trojic neničelnih koeficientov srednjih frekvenc
	triplets(command, block)

}

func triplets(command Command, block []float32) []float32 {
	//4-32, bigger the thr, smaller the span
	span := Span{0, 0}
	span.start = 4
	span.end = 32
	if command.thr > 32 {
		span.end = 64 - int(command.thr)
	}
	rand.Seed(int64(command.seed))
	//Keep track of what we picked
	picked := make([]int, 0)
	//Pick an M amount of triplets
	for i := 0; i < int(command.triplets)*3; i += 3 {
		random := rng(span)
		for j := 0; j < len(picked); j++ {
			if picked[j] == random || picked[j] == random+1 || picked[j] == random+2 {
				random = rng(span)
				j = 0
			}
		}
		picked = append(picked, random)
		picked = append(picked, random+1)
		picked = append(picked, random+2)

	}
	fmt.Println(picked)
	pickedtriplets := make([]float32, len(picked))
	for z := range picked {
		pickedtriplets[z] = block[picked[z]]
	}

	return pickedtriplets

}



func rng(span Span) int {
	return rand.Intn((span.end - span.start + 1)) + span.start
}
