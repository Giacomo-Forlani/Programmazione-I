package main

import "fmt"

// nomeCarta converte un numero intero in una stringa rappresentativa di una carta da gioco.
// Le carte sono associate a semi ("♠", "♥", "♣", "♦") e valori (1-10, J, Q, K).
func nomeCarta(x int) string {
	var seme, val string
	switch x / 13 {
	case 0:
		seme = "♠" // Picche
	case 1:
		seme = "♥" // Cuori
	case 2:
		seme = "♣" // Fiori
	case 3:
		seme = "♦" // Quadri
	}
	switch x % 13 {
	case 0: val = "2"
	case 1: val = "3"
	case 2: val = "4"
	case 3: val = "5"
	case 4: val = "6"
	case 5: val = "7"
	case 6: val = "8"
	case 7: val = "9"
	case 8: val = "10"
	case 9: val = "J"
	case 10: val = "Q"
	case 11: val = "K"
	case 12: val = "A"
	}
	return val + seme
}

func main() {
	// Esempio di utilizzo: stampa il nome di una carta per un valore dato.
	var x int
	fmt.Print("Inserisci un valore intero da 0 a 51: ")
	fmt.Scan(&x)

	// Stampa il nome della carta.
	fmt.Println("La carta è:", nomeCarta(x))
}