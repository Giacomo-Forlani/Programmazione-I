package main

import "fmt"

func main() {

	fmt.Print("Estrazione della parte decimale\n")
	var f, g float64
	fmt.Print("n: ")
	fmt.Scan(&f)
	g = f - float64(int(f))
	fmt.Print(g)
}