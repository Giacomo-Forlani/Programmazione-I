package main
import "fmt"

func main(){
	
	var n,i int
	fmt.Print("Qante volte vuoi essere salutato, Padrone?")

	fmt.Scan(&n)
	for i = 0; i < n ; i++ {
		fmt.Println ("ciao!")
	}
}