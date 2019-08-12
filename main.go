package main

import "infoObfuscation/input"

//Read bitmap
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

func main() {
	//Parse command line arguments
	command := input.Read()

	
}
