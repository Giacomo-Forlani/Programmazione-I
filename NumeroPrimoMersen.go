package main

import (
	"fmt"
	"math"
)

// Funzione per controllare se un numero è primo
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for d := 2; d <= int(math.Sqrt(float64(n))); d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

// Funzione per controllare se un numero è un numero primo di Mersenne
func isMersennePrime(x int) bool {
	if !isPrime(x) {
		return false
	}
	// Controlla se x è della forma 2^p - 1
	p := 1
	for {
		mersenne := int(math.Pow(2, float64(p))) - 1
		if mersenne == x {
			return true
		} else if mersenne > x {
			return false
		}
		p++
	}
}

func main() {
	fmt.Print("Controllo se è un numero primo di Mersenne. Inserisci un numero: ")
	var x int
	fmt.Scan(&x)

	if isMersennePrime(x) {
		fmt.Println("È un numero primo di Mersenne.")
	} else {
		fmt.Println("Non è un numero primo di Mersenne.")
	}
}
