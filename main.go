package main

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"os"
	"unsafe"
)

/*

vaja3 <vhodna slika> <opcija> <vhodno/izhodno sporočilo> <THR> <M>

kjer:

    <vhodna datoteka> - pot do poljubne BMP datoteke, ki vsebuje sivinsko sliko.
    <opcija>:
        h - skrivanje sporočila
        e - ekstrakcija sporočila
    <vhodno/izhodno sporočilo> - pot do vhodnega/izhodnega tekstovnega sporočila.
    <THR> - prag pri kompresiji
    <M> - število unikatnih množic trojic neničelnih koeficientov, ki se uporabijo v F5 steganografiji.


*/
//Bitmap stuff, x,y -> dimensions

type Bitmap struct {
	model color.Model
	x, y  image.Rectangle
}
type BitmapDimensions struct {
	x, y int
}

func main() {
	//Parse command line arguments

	command := read()

	rand.Seed(int64(command.seed))

	if command.option == "h" {

		pixels, bitmapdim := bitmapr(command.bmp)

		text := textr(command.message)

		textbits := text2bits(text)

		zigzag := eightxeight(pixels)

		quantized := make([][][]uint32, len(zigzag))

		for i := range zigzag {
			quantized[i] = quantize(command.thr, zigzag[i])
		}

		candidates := make([]int, 0)
		candidateblocks := make([][][]uint32, 0)

		wc := 0

		for i := 0; i < len(textbits)/int((command.tripletsnum * 2)); i++ {
			x := generaterng(candidates,len(quantized))

			candidates = append(candidates, x)
			words := textbits[wc : wc+int(command.tripletsnum*2)]

			candidateblocks = append(candidateblocks, f5(command, flatten(quantized[x]), words))
			wc += int(command.tripletsnum * 2)

		}

		stegblocks := blockarize(quantized, candidateblocks, candidates)

		filew("bitmaps/output", serialize(stegblocks, bitmapdim))
	} else {

		fmt.Println("Decoding", command.bmp)

		readb := filer(command.bmp)

		deserialized := deserialize(readb)
		dim := deserialized[0]
		fmt.Println("Dimension of coefficients", dim)

		//Cut out the dimensions
		deserialized = deserialized[1:]

		reconstructed := reconstruct3D(deserialized)

		message := inversionF5Caller(command, deserialized, reconstructed)
		extracted := make([]byte, 0)
		for i := 0; i < len(message); i += 8 {
			extracted = append(extracted, bitSliceToByte(message[i:i+8]))
		}
		fmt.Println(extracted)
		filew(command.message, extracted)
		extr := string(extracted)
		fmt.Println("Message:", extr)
		fmt.Println("Meta analysis: ")
		inverseHaar(command,reconstructed,dim)

	}
}

func generaterng(candidates []int,end int) int {
	x := rng(Span{0, end - 1})
	//Check if it's already generated
	for j := 0; j < len(candidates);j++ {
		if x == candidates[j] {

			x = rng(Span{0, end-1})
			j = -1
		}
	}
	return x
}

func reconstruct3D(deserialized []uint32) [][][]uint32 {
	reconstructed := make([][][]uint32, 0)
	for i := 0; i < len(deserialized)-64; i += 64 {
		recon := reconstructuint(deserialized[i : i+64])
		if recon == nil {
			fmt.Println("Unable to deserialize file")
			os.Exit(1)
		}
		reconstructed = append(reconstructed, recon)

	}
	return reconstructed
}

func reconstruct3DFloat(deserialized []float32) [][][]float32 {
	reconstructed := make([][][]float32, 0)
	for i := 0; i < len(deserialized)-64; i += 64 {
		recon := reconstructfloat(deserialized[i : i+64])
		if recon == nil {
			fmt.Println("Unable to deserialize file")
			os.Exit(1)
		}
		reconstructed = append(reconstructed, recon)

	}
	return reconstructed
}

func deserialize(bytesd []byte) []uint32 {
	des := make([]uint32, len(bytesd)/4)
	c := 0
	for i := 0; i < len(bytesd); i += 4 {
		des[c] = binary.LittleEndian.Uint32(bytesd[i : i+4])
		c++
	}
	return des
}

//Get the f5'd blocks into the thing
func blockarize(quantized [][][]uint32, candidateblocks [][][]uint32, candidates []int) [][][]uint32 {

	for i := 0; i < len(candidates);i++ {
		currcan := candidates[i]
		quantized[currcan] = candidateblocks[i]

	}
	fmt.Println(candidates)
	return quantized
}
func serialize(quantized [][][]uint32, dimensions BitmapDimensions) []byte {
	serialized := make([]byte, 0)
	//Encode the dimensions NxN at the start

	b := (*[4]byte)(unsafe.Pointer(&dimensions.x))[:]
	serialized = append(serialized, b...)
	for i := 0; i < len(quantized); i++ {
		for j := range quantized[0] {
			for l := range quantized[0][0] {
				curr := quantized[i][j][l]
				a := (*[4]byte)(unsafe.Pointer(&curr))[:]
				serialized = append(serialized, a...)
			}
		}
	}

	return serialized
}

/*

1. Tekstovno sporočilo, ki ga boste skrili najprej binarizirate. Na začetek binariziranega sporočila dodajte 4 zloge, kjer s celim številom (uint) poveste velikost sporočila v bitih.
2. Sivinsko sliko razdelite na 8x8 bloke pikslov.
3. Posamezni blok pretvorite v frekvenčni prostor na podlagi DCT ali Haarove transformacije (po želji).
4. Izvedite kvantizacijo frekvenčnega prostora v bloku, kjer najprej serializirate 64 koeficientov z algoritmom cik-cak nato pa postavite zadnjih THR koeficentov na 0 (uporabnik določi THR). Koeficiente tudi zaokrožite v cela števila (integer).
5. Izberite naključni blok (nad katerim še niste uporabili F5) ter nad (64-THR) koeficienti bloka izvedite algoritem F5:
		Izberete M naključnih unikatnih trojic neničelnih koeficientov srednjih frekvenc (indeksi od 4 do 32, v primeru večjega THR se ta razpon zmanjša!). Trojice med sabo nimajo prekrivanja.
        Vsako trojico določajo koeficienti AC1, AC2 in AC3, pri čemer korespodenčni LSB biti danih koeficentov so definirani kot C1=LSB(AC1), C2=LSB(AC2) in C3=LSB(AC3).
        Za vsako trojico (C1,C2,C3) vzamite 2 bita binariziranega sporočila, definirana kot x1 in x2 in izvedite naslednje operacije za skrivanje x1 in x2:
 6. Izbira po želji: a) direktno hranite v binarno datoteko bloke koeficientov. b) Bloke nazaj pretvorite v domeno pikslov ter po pretvorbi shranite rezultat kot izhodno sliko  BMP. V tem primeru bodite pozorni na zaokrožitvene napake pri izbiri AC-jev pred skrivanjem sporočila.
*/
