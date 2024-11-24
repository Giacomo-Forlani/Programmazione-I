// Questo programma calcola il massimo comune divisore (MCD) di due numeri interi positivi
// utilizzando l'algoritmo di Euclide, che è un metodo molto efficiente basato sulla divisione con resto.

package main

import "fmt"

func main() {
    var x, y, r int // Dichiarazione delle variabili: x e y sono gli input, r è il resto

    fmt.Println("Inserisci due numeri interi positivi separati da uno spazio:")
    fmt.Scan(&x, &y) // Legge i valori di x e y inseriti dall'utente

    // Calcolo dell'MCD utilizzando l'algoritmo di Euclide
    r = x % y // Calcola il resto della divisione di x per y
    for r != 0 { // Continua finché il resto non è zero
        x = y    // Aggiorna x al valore di y
        y = r    // Aggiorna y al valore del resto
        r = x % y // Calcola il nuovo resto
    }

    // Quando il resto diventa zero, y contiene l'MCD
    fmt.Println("Il massimo comune divisore è ", y)
}

