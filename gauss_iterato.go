package main

import "fmt"

// Il programma calcola la somma dei numeri interi da 1 a un numero 'n' fornito in input.
// Utilizza un approccio iterativo con un ciclo 'for'.

func main() {
    var n, c, s int // Dichiarazione delle variabili: 'n' è il limite superiore, 'c' è un contatore, 's' è la somma.
    
    fmt.Scan(&n)    // Legge un numero intero da input e lo assegna a 'n'.

    c = 1           // Inizializza il contatore 'c' a 1.
    for c <= n {    // Ciclo che esegue fino a quando 'c' è minore o uguale a 'n'.
        s += c      // Somma il valore di 'c' alla variabile 's'.
        c++         // Incrementa il contatore 'c' di 1.
    }

    fmt.Println(s)  // Stampa la somma calcolata.
}
