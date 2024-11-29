package main

import "fmt"

// calcolaCoefficienteAngolare calcola il coefficiente angolare della retta
// passante per i punti (x1, y1) e (x2, y2).
func calcolaCoefficienteAngolare(x1, y1, x2, y2 float64) (float64, error) {
    if x1 == x2 {
        return 0, fmt.Errorf("retta verticale, coefficiente angolare non definito")
    }
    return (y2 - y1) / (x2 - x1), nil
}

func main() {
    var x1, y1, x2, y2 float64
    fmt.Print("Inserisci le coordinate di due punti (x1 y1 x2 y2): ")
    fmt.Scan(&x1, &y1, &x2, &y2)

    coeff, err := calcolaCoefficienteAngolare(x1, y1, x2, y2)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Il coefficiente angolare della retta è: %.2f\n", coeff)
    }
}
