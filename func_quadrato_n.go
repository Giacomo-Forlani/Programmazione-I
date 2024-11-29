package main

import (
	"fmt"
)

// sqr calcola il quadrato di un valore intero ricevuto come parametro
func sqr(n int) int {
	// Calcolo del quadrato
	quadrato := n * n
	return quadrato
}

func main() {
	// Dichiarazione della variabile per l'input
	var n int

	// Richiesta dell'input all'utente
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&n)

	// Calcolo del quadrato del numero usando la funzione sqr
	result := sqr(n)

	// Stampa del risultato
	fmt.Println("Il quadrato di", n, "è", result)
}