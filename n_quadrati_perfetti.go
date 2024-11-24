package main

import "fmt"

// Questo programma legge un numero intero 'n' da input e calcola il quadrato di ogni numero
// intero compreso tra 0 e n-1, stampando ciascun risultato su una nuova riga.

func main() {
    var n int // Dichiarazione della variabile 'n', che rappresenta il limite superiore (esclusivo).

    fmt.Scan(&n) // Legge un numero intero da input e lo assegna alla variabile 'n'.

    for i := 0; i < n; i++ { // Ciclo che va da 0 a n-1.
        fmt.Println(i * i)   // Calcola il quadrato del numero corrente 'i' e lo stampa.
    }
}
