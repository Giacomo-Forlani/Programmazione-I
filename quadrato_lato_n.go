package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		for i := 0; i < n; i++{
		fmt.Print("* ")
		}
		fmt.Println()
	}
}