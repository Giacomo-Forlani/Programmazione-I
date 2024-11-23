package main
import "fmt"
func main() {
	var n, p int
	fmt.Scan(&n)
	for i:=0; i<=10; i++{
		p = i * n
		fmt.Print(p, "\t")
	}
}
