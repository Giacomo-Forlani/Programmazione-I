package main

import (
	"fmt"
)

func main(){
	var n int
	fmt.Print("Fino che numero di Fibonacci vuoi stmpare?: ")
	fmt.Scan(&n)
	var p, s int = 1, 1
	fmt.Println("*")
	fmt.Println("*")
	for i := 2; i < n; i++{
		if i % 2 == 0{
			p = s + p
			for j := 0; j < p; j++{
				fmt.Print("*")
			}
			fmt.Print(p)
		}else{
			s = s + p
			for j := 0; j < s; j++{
				fmt.Print("*")
			}
			fmt.Print(s)
		}
		fmt.Println("")
	}
}