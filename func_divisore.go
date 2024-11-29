package main

import "fmt"

// isDivisor verifica se un numero `a` è divisore di `b`.
func isDivisor(a, b int) bool {
    return b%a == 0
}

func main() {
    var a, b int
    fmt.Print("Inserisci due numeri (a b): ")
    fmt.Scan(&a, &b)
    if isDivisor(a, b) {
        fmt.Printf("%d è divisore di %d\n", a, b)
    } else {
        fmt.Printf("%d non è divisore di %d\n", a, b)
    }
}
