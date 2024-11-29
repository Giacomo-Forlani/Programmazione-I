package main

import (
    "fmt"
    "strconv"
)

// countZeros conta quanti zeri ci sono nella rappresentazione di un numero.
func countZeros(n int) int {
    count := 0
    str := strconv.Itoa(n) // Converte il numero in stringa
    for _, char := range str {
        if char == '0' {
            count++
        }
    }
    return count
}

func main() {
    totalZeros := 0
    for i := 1; i <= 1_000_000; i++ {
        totalZeros += countZeros(i) // Conta gli zeri per ogni numero
    }
    fmt.Printf("Il numero totale di zeri da 1 a 1.000.000 è: %d\n", totalZeros)
}
