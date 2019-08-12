package main

import (
	"fmt"
	"infoObfuscation/binaryIo"
	"infoObfuscation/input"
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

func main() {
	//Parse command line arguments
	command := input.Read()
	pixels := binaryIo.Bitmapr(command.Bmp)
	text := binaryIo.Textr(command.Message)
	fmt.Println(len(pixels) , "x", len(pixels[0]), " TEXT : " , len(text))

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
