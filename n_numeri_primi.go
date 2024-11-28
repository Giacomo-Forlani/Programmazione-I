package main

import	"fmt"

func main(){
	var n, d, c int
	fmt.Scan(&n)
	for x := 2; c < n; x++{
		if x % d == 0{
			break
		}
	}
	if d == x{
		fmt.Println(x)
		c++
		if c == n{
			break
		}
	}
}