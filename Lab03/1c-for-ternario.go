package main

import "fmt"

func main() {
    /*
    Programma per sommare 5 numeri inseriti dall'utente
    */
    const K = 5
	var n int
	somma := 0
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		somma += n  // somma = somma + n
	}
	fmt.Println(somma)
}
