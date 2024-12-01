package main

import "fmt"

/*

     ********
      *     * 0
       *    * 1
        *   * 2
         *  *
          * *
           **
            *

  Dato n stampa il triangolo di base e altezza n visualizzato sopra (n=7). */


func main() {
	var n int
	fmt.Scan(&n)

	// Prima riga
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	// Righe intermedie
	for r := 0; r < n - 2; r++ {
		for i := 0; i < r + 1; i++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		for i := 0 ; i < (n - 2) - (r + 1); i++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}

	// Ultima riga
	for i := 0; i < n - 1; i++ {
		fmt.Print(" ")
	}
	fmt.Println("*")

}