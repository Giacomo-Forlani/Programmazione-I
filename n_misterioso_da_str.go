package main

import (
	"fmt"
	"strconv"
)

func Misterioso(s string) (int, bool) {
	pulita := ""
	for _, r := range s {
		if r != '*' {
			pulita += string(r)
		}
	}
	num, err := strconv.Atoi(pulita)
	if err == nil {
		return num, true
	} else {
		return 0, false
	}
}

func main() {
	fmt.Println(Misterioso("123"))   // Expected: 123, true
	fmt.Println(Misterioso("1*2*3")) // Expected: 123, true
	fmt.Println(Misterioso("abc"))   // Expected: 0, false
	fmt.Println(Misterioso("12*3a")) // Expected: 0, false
}
