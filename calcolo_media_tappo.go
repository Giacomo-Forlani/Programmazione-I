package main

import "fmt"

// Il programma calcola la media aritmetica dei numeri interi inseriti dall'utente.
// L'inserimento termina quando l'utente inserisce '0'.

func main() {
    var n, h, s int // 'n' conta i numeri validi, 'h' memorizza il numero corrente, 's' accumula la somma.

    for {
        fmt.Scan(&h) // Legge un numero intero da input.
        if h == 0 {  // Se l'utente inserisce '0', interrompe il ciclo.
            break
        }
        s += h // Aggiunge il numero letto alla somma totale.
        n++    // Incrementa il contatore dei numeri validi.
    }

    // Calcola la media come rapporto tra somma e numero di elementi, con conversione in float64.
    media := float64(s) / float64(n)

    fmt.Println("Media:", media) // Stampa la media calcolata.
}
