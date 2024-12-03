package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const k = 11
	var n int
	for i:=1; i<=k; i++{
		n = rand.Intn(k)
		fmt.Print(n, " ")
	}
}