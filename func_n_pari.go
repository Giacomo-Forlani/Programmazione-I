package main

import "fmt"

func isEven(x int) bool{
	return(x%2==0)
}

func main(){
	var n int 
	fmt.Scan(&n)
	if isEven(n){
		fmt.Println("È Pari")
	}else{
		fmt.Println("È Dispari")
	}
}