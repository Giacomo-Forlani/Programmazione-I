package main

import "fmt"

func main() {
	fmt.Println("Media tra 2 altezze")
	var h1, h2 int
	var media float64
	fmt.Print("n1: ")
	fmt.Scan(&h1)
	fmt.Print("n2: ")
	fmt.Scan(&h2)
	media = float64(h1+h2) / 2
	fmt.Print(media)
}