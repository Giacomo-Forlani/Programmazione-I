package main

import "fmt"

func main() {
	var n,h,s int
	fmt.Print("Quante persone? ")
	fmt.Scan(&n)
	for i := 0; i<n; i++{
		fmt.Scan(&h)
		s += h
	}
	media := float64(s)/float64(n)
	fmt.Println(media)
}