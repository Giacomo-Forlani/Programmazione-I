package main

import "fmt"

func main(){
	var b, d, e int
	var c = 1
	a := []int{1,1,2,2,3,3,3,4,4,4,4,5,6,6,7,7,7,7,7,2}
	for i := 0; i <= len(a); i++{
		b = a[i]
		if  b == a[i + 1]{
			c = c + 1
		}
	
	}
}