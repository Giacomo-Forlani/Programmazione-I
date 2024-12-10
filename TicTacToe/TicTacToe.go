package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Crea una matrice bidimensionale di dimensione `dim x dim` e la restituisce.
func matrixCreation(dim int) [][]int {
    matrix := make([][]int, dim)
    for i := range matrix {
        matrix[i] = make([]int, dim)
    }
    return matrix
}

// Stampa la matrice con simboli specifici per i valori 0, 1 e 2.
func printMatrix(matrix [][]int) {
    for _, row := range matrix {
        for _, value := range row {
            if value == 1 {
                fmt.Print("⭕")
            } else if value == 2 {
                fmt.Print("❌")
            } else {
                fmt.Print("🧻")
            }
        }
        fmt.Println()
    }
    fmt.Println()
}

// Aggiunge un valore alla matrice nella posizione specificata se valida.
func addValue(matrix [][]int, dim, row, col, value int) ([][]int, bool) {
    var checker bool = true
    if row >= 0 && row < dim && col >= 0 && col < dim && matrix[row][col] == 0 {
        matrix[row][col] = value
        checker = false
    } else {
        fmt.Println("Valori non validi")
    }
    return matrix, checker
}

// Genera un numero casuale tra 0 e `dim-1`.
func randInt(dim int) int {
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)
    return r.Intn(dim)
}

// Calcola il numero di mosse necessarie e se il numero di mosse è dispari.
func calcMoves(dim int) (int, bool) {
    if dim%2 == 0 {
        return dim * dim / 2, false
    } else {
        return (dim * dim / 2) + 1, true
    }
}

func winner(matrix [][]int, dim int) bool{
	var referee bool = false
	for i := 0; i < dim; i++ {
		for j := 0; j < dim - 1; j++ {
            fmt.Printf("%d - %d\n", matrix[i][j], matrix[i+1][j+1])
			if matrix[i][j] != matrix [i + 1][j + 1]{
				break
			}
		}

	}
	return referee
}

func main() {
    var dim, moves int
    var cutter bool
    fmt.Print("Dimensione del campo: ")
    fmt.Scan(&dim)
    matrix := matrixCreation(dim)
    moves, cutter = calcMoves(dim)
    for i := 0; i < moves; i++ {
        var row, col, value int = 0, 0, 1
        var checker bool = true
        for checker {
            fmt.Print("Inserisci il numero della riga in cui vuoi posizionare la mossa: ")
            fmt.Scan(&row)
            fmt.Print("Inserisci il numero della colonna in cui vuoi posizionare la mossa: ")
            fmt.Scan(&col)
            matrix, checker = addValue(matrix, dim, row, col, value)
        }
        printMatrix(matrix)
        if i == moves-1 && cutter == true {
            break
        }
        checker = true
        value = 2
        for checker {
            row = randInt(dim)
            col = randInt(dim)
            matrix, checker = addValue(matrix, dim, row, col, value)
        }
        printMatrix(matrix)
    }
    fmt.Println("Fine della partita")
	referee := winner(matrix, dim)
	fmt.Print(referee)
}