package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		p:=1
		for c:=0;c<i;c++{
			p*=2
		}
		fmt.Println(p)
	}
}