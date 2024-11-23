package main

import "fmt"

func main() {
	fmt.Println("Stabilire se una data precede la seconda")
	var g1, m1, a1, g2, m2, a2 int
	fmt.Print("Giorno prima data: ")
	fmt.Scan(&g1)
	fmt.Print("Mese prima data: ")
	fmt.Scan(&m1)
	fmt.Print("Anno prima data: ")
	fmt.Scan(&a1)
	fmt.Print("giorno seconda data: ")
	fmt.Scan(&g2)
	fmt.Print("Mese seconda data: ")
	fmt.Scan(&m2)
	fmt.Print("anno seconda data: ")
	fmt.Scan(&a2)
	if a1 < a2 {
		fmt.Print("La prima data precede la seconda")
	} else if a1 > a2 {
		fmt.Print("La prima data non precede la seconda")
	} else {
		if m1 < m2 {
			fmt.Print("La prima data precede la seconda")
		} else if m1 > m2 {
			fmt.Print("La prima data non precede la seconda")

		} else {
			if g1 < g2 {
				fmt.Print("La prima data precede la seconda")
			} else if g1 > g2 {
				fmt.Print("La prima data non precede la seconda")
			}else{
				fmt.Print("Le date sono uguali")
			}
		}
	}
}