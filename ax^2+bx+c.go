// Questo programma risolve un'equazione quadratica del tipo ax^2 + bx + c = 0.
// Utilizza la formula risolutiva delle equazioni quadratiche e gestisce i casi 
// in cui il discriminante (delta) è positivo, zero o negativo.

package main

import (
	"fmt"  // Pacchetto per l'input/output
	"math" // Pacchetto per funzioni matematiche come radice quadrata
)

func main() {
	var a, b, c float64 // Dichiarazione delle variabili per i coefficienti a, b e c
	fmt.Println("Inserisci i coefficienti a, b e c separati da uno spazio:")
	fmt.Scan(&a, &b, &c) // Lettura dei coefficienti da input

	// Calcolo del discriminante (delta)
	delta := b*b - 4*a*c

	// Controllo se il discriminante è vicino a zero (tenendo conto dell'approssimazione numerica)
	if math.Abs(delta) < 1E-6 {
		// Caso in cui delta è quasi zero: una soluzione reale doppia
		fmt.Println("Probabilmente ci sono 2 soluzioni coincidenti x= ", -b / (2 * a))
	} else if delta > 0 {
		// Caso in cui delta è positivo: due soluzioni reali distinte
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println("Le 2 soluzioni reali distinte sono: x1= ", x1, "x2= ", x2)
	} else {
		// Caso in cui delta è negativo: nessuna soluzione reale
		fmt.Println("Non esistono soluzioni reali per i coefficienti forniti.")
	}
}
