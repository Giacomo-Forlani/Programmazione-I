package main

import "fmt"

func main(){
	var h,m int
	fmt.Scan(&h)
	fmt.Scan(&m)
	x := (24 - h) * 60 - m
	fmt.Println(x)
}