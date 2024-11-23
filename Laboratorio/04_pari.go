package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const k = 10
	var n, pari int
	for i:=1; i<=k; i++{
		n = rand.Intn(k)
		fmt.Print(n, " ")
		if n%2==0{
			pari++
		}
	}
	fmt.Print("\n", pari)
}