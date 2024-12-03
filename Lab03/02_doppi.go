package main

import "fmt"

func main(){
	const k = 5
	var n int
	for i:=0; i<=5; i++{
		fmt.Print("Digita un valore: ")
		fmt.Scan(&n)
		n = n * 2
		fmt.Print(n, "\n")
	}
}