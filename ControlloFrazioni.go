package main

import "fmt"

func main() {
	fmt.Print("Controllo se una frazione è minore di un'altra\n")
	var n1, n2, d1, d2 int
	fmt.Print("Numeratore primo valore: ")
	fmt.Scan(&n1)
	fmt.Print("Denominatore primo valore: ")
	fmt.Scan(&d1)
	fmt.Print("Numeratore secondo valore: ")
	fmt.Scan(&n2)
	fmt.Print("Denominatore secondo valore: ")
	fmt.Scan(&d2)
	if n1*d2 < d1*n2 {
		fmt.Print("La prima è minore")
	} else {
		fmt.Print("La seconda è maggiore o uguale")
	}
}