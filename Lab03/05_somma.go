package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const k = 11
	var n, somma int
	for i:=1; i<=k; i++{
		n = rand.Intn(k) + 1
		fmt.Print(n, " ")
		somma = n + somma
	}
	fmt.Print("\n", somma)
}