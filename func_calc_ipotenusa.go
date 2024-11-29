package main

import (
	"fmt"
	"math"
)

func ipotenusa(cat1, cat2 float64) float64{
	var x float64
	x = cat1*cat1 + cat2*cat2
	return math.Sqrt(x)
}

func main(){
	x := ipotenusa(7.0, 8.5)
	fmt.Println(x)
}