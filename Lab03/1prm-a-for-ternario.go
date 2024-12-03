package main
import "fmt"
func main() {
    /*
    Programma che stampa tutti i multipli di 2 minori o uguale del numero inserito dall'utente
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
