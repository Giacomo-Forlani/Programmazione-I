package main

import "fmt"

func main() {
	fmt.Print("Arrotondare un numero\n")
	var f float64
	var x int
	fmt.Print("n: ")
	fmt.Scan(&f)
	x = int(f + 0.5)
	fmt.Print(x)
}