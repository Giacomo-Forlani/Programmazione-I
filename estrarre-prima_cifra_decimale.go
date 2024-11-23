package main

import "fmt"

func main() {

	fmt.Print("Estrazione prima cifra decimale\n")
	var f float64
	var g int
	fmt.Print("n: ")
	fmt.Scan(&f)
	g = int(10 * (f - float64(int(f))))
	fmt.Print(g)
}