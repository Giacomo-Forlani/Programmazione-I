package main

/* Lo scopo di questo programma è calcolare la differenza esatta in giorni fra due date.
 * Il modo standard di rappresentare una data è mediante il numero di giorni trascorsi da una data fissa,
 * detta "epoca", che è il 1/1/1970.
 */

import "fmt"

/* Restituisce true sse a è un anno bisestile */
func bisestile(a int) bool {
	return (a % 4 == 0 && a % 100 != 0) || (a % 400 == 0)
}

/*
 * Restituisce il numero di giorni dal 1/1/1970 alla data g/m/a.
 */
func giorniDaEpoca(g, m, a int) int {
	somma := 0
	// Sommo gli anni interi fino ad a escluso
	for i := 1970; i < a; i++ {
		if bisestile(i) {
			somma += 366
		} else {
			somma += 365
		}
	}
	// Sommo i mesi interi dell'anno a da gennaio fino a m escluso 
	for i := 1; i < m; i++ {
		if i == 11 || i == 4 || i == 6 || i == 9 {
			somma += 30
		} else if i == 2 {
			if bisestile(a) {
				somma += 29
			} else {
				somma += 28
			}
		} else {
			somma += 31
		}
	} 
	somma += g 
	return somma - 1
}

func main() {
	var g1, m1, a1, g2, m2, a2 int
	fmt.Scan(&g1, &m1, &a1)
	fmt.Scan(&g2, &m2, &a2)
	x1 := giorniDaEpoca(g1, m1, a1)
	x2 := giorniDaEpoca(g2, m2, a2)
	fmt.Print("Dal ", g1, "/", m1, "/", a1, " al ", g2, "/", m2, "/", a2, " sono passati ", x2-x1, " giorni.\n")
}


// PER CASA: estendere a date precedenti all'epoca!


