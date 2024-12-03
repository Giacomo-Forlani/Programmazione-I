package main

import (
	"fmt"
)

func main() {
	var n, s int
	fmt.Print("Dimensioni della V: ")
	fmt.Scan(&n)
	s = n
	for i := 1; i <= n; i++{
		for j := 1; j < i; j++{
			fmt.Print("_")
		}
		fmt.Print("*")
		for j := 0; j <= s; j++{
			fmt.Print(j)
		}
		s -=2
		if i == n{
			break
		}
		fmt.Println("*")
	}
}