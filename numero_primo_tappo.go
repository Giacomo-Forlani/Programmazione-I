package main

import	"fmt"

func main(){
	var entrato bool
	var n int

	fmt.Scan(&n)
	for d:= 2; d<n; d++{
		if n%d==0{
			entrato = true
			break
		}
	}
	if entrato{
		fmt.Println("Composto")
	}else{
		fmt.Println("Primo")
	}
}