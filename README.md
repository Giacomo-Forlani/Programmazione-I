# Corso di Programmazio I dell'UniMi

## Comandi di base del compilatore
- `go build <NOME-FILE.go>`
- `./<NOME-FILE>`
> Per Windows: `<NOME-FILE>`
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
- 
### Dichiarazioni di variabile
- Dichiarazione di variabile singola: `var <VARIABILE> <TIPO>`
- Dichiarazione di variabile multipla `var <VARIABILE>,<VARIABILE>,<VARIABILE> <TIPO>`
### Assegnamento
- `<VARIABILE> = <ESPRESSIONE>`
- Short assignment: `<VARIABILE> := <ESPRESSIONE>`
> In questo caso non è necesario dichiarare prima la funzione
### Shadowing
> ⚠️ Le variabili dichiarate esistono solo all'interno della funzione 
> in cui sono state dichiarate e non in quele esterne o in quelle contenute

## Operatori
-