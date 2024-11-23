// Scrivere un programma che dato il prezzo di un bene e data l'aliquota IVA, determina l'imponibile

package main

import (
	"fmt"
)

func main() {
	var prezzo, aliquota float64

	// Input prezzo del bene
	fmt.Print("Inserisci il prezzo del bene (IVA inclusa): ")
	fmt.Scan(&prezzo)

	// Input aliquota IVA in percentuale
	fmt.Print("Inserisci l'aliquota IVA (%): ")
	fmt.Scan(&aliquota)

	// Calcolo dell'imponibile
	imponibile := prezzo / (1 + aliquota/100)

	// Stampa del risultato
	fmt.Println("L'imponibile del bene è: ", imponibile)
}
