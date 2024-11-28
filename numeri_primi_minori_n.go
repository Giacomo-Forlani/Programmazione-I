package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n) // Legge il numero fino a cui calcolare i numeri primi

    for x := 2; x <= n; x++ { // Ciclo esterno: controlla tutti i numeri da 2 a n
        var d int
        for d = 2; d < x; d++ { // Ciclo interno: verifica se x è divisibile per qualche numero minore di x
            if x%d == 0 { // Se x è divisibile per d, non è un numero primo
                break
            }
        }
        if d == x { // Se il ciclo interno termina senza trovare divisori, x è primo
            fmt.Println(x) // Stampa il numero primo
        }
    }
}
