package main

import "fmt"

func main(){
	var a, b, c int
	fmt.Print("Quale numero vuoi convertire?: ")
	fmt.Scan(&a)
	fmt.Print("In che base vuoi trasformare il numero?: ")
	fmt.Scan(&b)
	for a != 0{
		c = a % b
			if c != 0{
				fmt.Print("1")
			}else{
				fmt.Print("0")
			}
		a = a / b
	}
}