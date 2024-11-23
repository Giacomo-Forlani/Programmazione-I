# Corso di Programmazio I dell'UniMi

## Comandi di base del compilatore
- `go build <NOME-FILE.go>`
- `./<NOME-FILE>`
> Per Windows solo: `<NOME-FILE>`
- `go run <NOME-FILE.go>`
- `go doc fmt.Print`
- `go help`
- `go fmt <NOME-FILE>`

## Commenti in GO
- `// ...`
- `/* ... */`

## Variabili
### Tipi di Variabile
- int
- float64
-

### Dichiarazioni di variabile
- Dichiarazione di variabile singola: `var <VARIABILE> <TIPO>`
- Dichiarazione di variabile multipla `var <VARIABILE>,<VARIABILE>,<VARIABILE> <TIPO>`
### Assegnamento
- `<VARIABILE> = <ESPRESSIONE>`
- Short assignment: `<VARIABILE> := <ESPRESSIONE>`
> Nelle short assignment non è necesario dichiarare prima la funzione
### Shadowing❗️
> ⚠️ Le variabili dichiarate esistono solo all'interno della funzione in cui sono state dichiarate e non in quelle esterne o in quelle contenute

## Operatori
- Somma: `+`
- Sottrazione: `-`
- Prodotto: `*`
- Divisione: `/`
- Resto: `%`

## I/O di base
### Input
- `fmt.Scan`
    - `fmt.Scan(&a)`
### Output
- Stampare sulla riga corrente: `fmt.Print`
    - `fmt.Print("Risultati", a, "e", b)`
    - `fmt.Print(a, b, a + b / a)`
- Stampare a capo: `fmt.Println`
    - `fmt.Println("Risultati", a, "e", b)`
    - `fmt.Println(a, b, a + b / a)`

> Esercizio: Scrivere un programma che dato il prezzo di un bene e data l'aliquota IVA, determina l'imponibile


