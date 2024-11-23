package main
import "fmt"
func main() {
    /*
		Programma per stampare il simbolo "*" il numero di volte deciso dall'utente
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
