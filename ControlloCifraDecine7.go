package main

import "fmt"

func main() {
if selezione == 6{
	fmt.Print("Stabilire se la somma di 2 interi ha la cifra delle decine pari a 7")
	var a, b, c int
	fmt.Print("Primo valore: ")
	fmt.Scan(&a)
	fmt.Print("Seconda valore: ")
	fmt.Scan(&b)
	c = (a+b)/10
	if c == 7{
		fmt.Print("La cifra delle decine della somma è uguale a 7")
	}else{
		fmt.Print("La cifra delle decine della somma non è uguale a 7")
	}

}