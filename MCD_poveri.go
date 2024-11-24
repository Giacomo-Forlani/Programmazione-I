// Questo programma calcola il massimo comune divisore (MCD) di due numeri interi positivi x e y
// utilizzando un metodo iterativo basato sulla riduzione del divisore.

package main

import "fmt"

func main() {
    var x, y, m int // Dichiarazione delle variabili: x e y sono gli input, m sarà il massimo comune divisore (MCD)

    fmt.Println("Inserisci due numeri interi positivi separati da uno spazio:")
    fmt.Scan(&x, &y) // Legge i valori di x e y inseriti dall'utente

    // Determina il valore iniziale di m come il minore tra x e y
    if x < y {
        m = x
    } else {
        m = y
    }

    // Ciclo per trovare il MCD
    // Continua a ridurre m finché non soddisfa la condizione di essere divisore sia di x che di y
    for x%m != 0 || y%m != 0 {
        m-- // Riduce m di 1 ad ogni iterazione
    }

    // Stampa il risultato (il valore di m dopo il loop è il MCD)
    fmt.Printf("Il massimo comune divisore di %d e %d è %d\n", x, y, m)
}
