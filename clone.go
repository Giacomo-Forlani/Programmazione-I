package main

import "fmt"

func cerca(x, y []int) []int {
	var z []int
	for _, ex := range x {
		var i int
		for i = 0; i < len(y); i++ {
			if y[i] == ex {
				break
			}
		}
		if i < len(y) {
			z = append(z, ex)
		}
	}
	return z
}

func main() {
	x := []int{2, 5, 7, 8, 4}
	y := []int{7, 3, 4, 9, 2}
	result := cerca(x, y)
	fmt.Println("Elementi comuni:", result)
}
