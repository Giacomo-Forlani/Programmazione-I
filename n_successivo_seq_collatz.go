package main

import "fmt"

/*
 * Dato n stampa la sequenza di Collatz
 * 
 * Sequenza di Collatz
 *
 * n -> n/2 se n è pari
 * n -> 3n+1 se n è dispari
 * quando arrivo a 1 smetto
 * 
 * 7 -> 22 -> 11 -> 34 -> 17 -> 52 -> 26 -> 13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1
 */

/* Dato n dice qual è il prossimo numero della sequenza */
func prossimo(n int) int {
	if n % 2 == 0 {
		return n / 2
	} else {
		return 3 * n + 1
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	for n > 1 {
		fmt.Print(n, " ")
		n = prossimo(n)
	}
	fmt.Println(1)
}

// PER CASA: Scrivere un programma che dato n stampa i primi n numeri naturali e accanto a ciascuno di essi
// un istogramma con tanti asterischi quanto è lunga la sequenza di Collatz a partire da quel numero





