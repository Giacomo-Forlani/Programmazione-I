package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Legge una stringa da standard input
	var input string
	fmt.Print("Inserisci una stringa: ")
	fmt.Scanln(&input)

	// Variabili per tracciare la presenza di maiuscole e minuscole
	hasLower := false
	hasUpper := false

	// Controlla ogni rune della stringa
	for _, r := range input {
		if unicode.IsLower(r) {
			hasLower = true
		}
		if unicode.IsUpper(r) {
			hasUpper = true
		}
	}

	// Stampa il risultato
	if hasLower && hasUpper {
		fmt.Println("sia minuscole che maiuscole")
	} else if hasLower {
		fmt.Println("solo minuscole")
	} else if hasUpper {
		fmt.Println("solo maiuscole")
	} else {
		fmt.Println("nessuna lettera minuscola o maiuscola")
	}
}
