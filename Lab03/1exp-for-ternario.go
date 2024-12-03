package main

import "fmt"

func main() {
    /*
    Programma che per visualizzare come i numeri possono essere rappresentati in base 8 mostrando tutte le possibili combinazioni di posizioni e valori fino al numero di input
    */

	const BASE = 8
	var n int

	fmt.Print("un int: ")
	fmt.Scan(&n)

	length := float64(n)
	for i := 0; float64(i)/BASE < length; i++ { // i < n*BASE
		fmt.Print("|\t")
	}
	fmt.Println()

	for i := 0; float64(i)/BASE < length; i++ {
		fmt.Print(i/BASE, i%BASE, "\t")
	}
	fmt.Println()
}
