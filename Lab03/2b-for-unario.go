package main

import "fmt"

func main() {
	var c, n, o int
	fmt.Scan(&n)
	for c != 100 {
		o = n
		fmt.Scan(&n)
		c = (o%100)*100 + n%100
		fmt.Println(c)
	}
}