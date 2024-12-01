package main

import (
	"fmt"
	"strconv"
)

func nomeCarta(x int) string {
	var val string
	if x%13 <= 8 {
		val = strconv.Itoa((x%13) + 2)
	} else {
		switch x % 13 {
		case 9:
			val = "J"
		case 10:
			val = "Q"
		case 11:
			val = "K"
		case 12:
			val = "A"
		}
	}
	return val
}

func main() {
	// Esempio di utilizzo: stampa il nome di una carta per un valore dato.
	var x int
	fmt.Print("Inserisci un valore intero da 0 a 12: ")
	fmt.Scan(&x)

	// Stampa il nome della carta.
	fmt.Println("La carta è:", nomeCarta(x))
}