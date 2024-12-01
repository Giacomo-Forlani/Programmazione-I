package main

import "fmt"

// inLettere converte un numero intero tra 1 e 99 nella sua rappresentazione in lettere in italiano.
func inLettere(x int) string {
    // Gestione dei numeri tra 10 e 19, che hanno forme specifiche.
    if x >= 10 && x <= 19 {
        switch x {
        case 10:
            return "dieci"
        case 11:
            return "undici"
        case 12:
            return "dodici"
        case 13:
            return "tredici"
        case 14:
            return "quattordici"
        case 15:
            return "quindici"
        case 16:
            return "sedici"
        case 17:
            return "diciassette"
        case 18:
            return "diciotto"
        case 19:
            return "diciannove"
        }
    } else if x >= 1 && x <= 99 { // Gestione dei numeri da 1 a 99, esclusi quelli già trattati.
        var dec, uni string

        // Determina la parte delle decine.
        switch x / 10 {
        case 2:
            dec = "venti"
        case 3:
            dec = "trenta"
        case 4:
            dec = "quaranta"
        case 5:
            dec = "cinquanta"
        case 6:
            dec = "sessanta"
        case 7:
            dec = "settanta"
        case 8:
            dec = "ottanta"
        case 9:
            dec = "novanta"
        }

        // Determina la parte delle unità.
        switch x % 10 {
        case 1:
            uni = "uno"
        case 2:
            uni = "due"
        case 3:
            uni = "tre"
        case 4:
            uni = "quattro"
        case 5:
            uni = "cinque"
        case 6:
            uni = "sei"
        case 7:
            uni = "sette"
        case 8:
            uni = "otto"
        case 9:
            uni = "nove"
        }

        // Elide l'ultima vocale delle decine quando segue "uno" o "otto".
        if x%10 == 1 || x%10 == 8 {
            dec = dec[:len(dec)-1]
        }

        return dec + uni // Combina le decine con le unità.
    }

    // Caso non valido (fuori dall'intervallo previsto).
    return "Numero non supportato"
}

func main() {
    var x int

    // Richiede un numero da 1 a 99.
    fmt.Print("Inserisci un numero intero da 1 a 99: ")
    fmt.Scan(&x)

    // Stampa la rappresentazione in lettere del numero.
    fmt.Println(inLettere(x))
}
