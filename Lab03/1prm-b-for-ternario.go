package main
import "fmt"
func main() {
    /*
    Programma che sottrae 1 al numero inserito dall'utente e lo stampa fino a quando non arriva a 0
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
