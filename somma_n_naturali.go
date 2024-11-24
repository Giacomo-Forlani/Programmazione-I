package main

import "fmt"

// Il programma calcola la somma dei numeri interi da 1 a un numero 'n' fornito in input.
// Successivamente confronta il risultato con la formula matematica di Gauss per verificare la correttezza del calcolo.

func main() {
    var n, c, s int // Dichiarazione delle variabili intere: 'n' è il numero massimo, 'c' è un contatore, 's' è la somma.
    fmt.Scan(&n)    // Acquisisce un numero intero da input e lo assegna a 'n'.
    
    c = 1           // Inizializza il contatore 'c' a 1.
    for c <= n {    // Ciclo che itera da 1 fino a 'n'.
        s += c      // Aggiunge il valore di 'c' alla somma 's'.
        c++         // Incrementa il contatore di 1.
    }
    fmt.Println(s)  // Stampa il valore della somma calcolata.

    gauss := n * (n + 1) / 2 // Utilizza la formula di Gauss per calcolare direttamente la somma.
    if s != gauss {          // Confronta il risultato calcolato con quello della formula.
        fmt.Println("***Panico") // Se i risultati non coincidono, stampa un messaggio di errore.
    }
}
