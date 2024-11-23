package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Funzione per ottenere la scelta del giocatore
func playerChoice() string {
	var choice string
	fmt.Println("Scegli tra Carta, Forbici, Sasso:")
	fmt.Scanln(&choice)
	choice = strings.ToLower(choice) // Converte la scelta in minuscolo per evitare problemi di maiuscole
	return choice
}

// Funzione per ottenere la scelta del computer
func computerChoice() string {
	choices := []string{"carta", "forbici", "sasso"}
	rand.Seed(time.Now().UnixNano()) // Inizializza il generatore di numeri casuali
	return choices[rand.Intn(len(choices))]
}

// Funzione per determinare il vincitore
func determineWinner(player, computer string) string {
	if player == computer {
		return "Pareggio!"
	}

	switch player {
	case "carta":
		if computer == "sasso" {
			return "Hai vinto! Carta batte Sasso."
		} else {
			return "Hai perso! Forbici batte Carta."
		}
	case "forbici":
		if computer == "carta" {
			return "Hai vinto! Forbici batte Carta."
		} else {
			return "Hai perso! Sasso batte Forbici."
		}
	case "sasso":
		if computer == "forbici" {
			return "Hai vinto! Sasso batte Forbici."
		} else {
			return "Hai perso! Carta batte Sasso."
		}
	default:
		return "Scelta non valida!"
	}
}

func main() {
	// Ottieni le scelte
	player := playerChoice()
	computer := computerChoice()

	fmt.Printf("Il computer ha scelto: %s\n", computer)

	// Determina e mostra il vincitore
	result := determineWinner(player, computer)
	fmt.Println(result)
}
