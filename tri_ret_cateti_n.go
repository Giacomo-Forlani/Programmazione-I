package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		for c := 0; c<i+1; c++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}