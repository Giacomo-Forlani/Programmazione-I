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
- bool
### Dichiarazioni di variabile
- Dichiarazione di variabile singola: `var <VARIABILE> <TIPO>`
- Dichiarazione di variabile multipla `var <VARIABILE>,<VARIABILE>,<VARIABILE> <TIPO>`
### Shadowing
> ⚠️ Le variabili dichiarate esistono solo all'interno della funzione in cui sono state dichiarate e non in quelle esterne o in quelle contenute
### Mescolare int e float64

#### Esercizi
- (`arrotondamento.go`)
- (`estrazione_parte_frazioanria.go`)
- (`estrarre-prima_cifra_decimale.go`)
- (`media_tra_altezze.go`)
### Blank variable
- `_ = <ESPRESSIONE>`
- `_ = <VARIABILE>`
> La blank variable serve ad assegnare cose che non hai intenzione di usare, come fosse un cestino, ma senza la necessità di cancellare la variabile

## Operatori
### int
- Somma: `+`
- Sottrazione: `-`
- Prodotto: `*`
- Divisione: `/`
- Resto: `%`
### float
- Somma: `+`
- Sottrazione: `-`
- Prodotto: `*`
- Divisione: `/`
### Incremento/Decremento
- Incremento di 1: `x++`
- Decremento di 1: `x--`
### Operatori di assegnamento
- Somma: `+=`
- Sottrazione: `-=`
- Prodotto: `*=`
- Divisione: `/=`
- Resto: `%=`
> `%=` valido solo con int
### Operatori di confronto
- Maggiore: `>`
- Maggiore Uguale: `>=`
- Minore: `<`
- Minore Uguale: `<=`
- Uguale: `==`
- Diverso: `!=`
### bool
- `&&`

| x     | y     | `x && y` |
|-------|-------|--------|
| true  | true  | true   |
| true  | false | false  |
| false | true  | false  |
| false | false | false  |

- `||`

| x     | y     | x \|\|t y |
|-------|-------|--------|
| true  | true  | true   |
| true  | false | true   |
| false | true  | true   |
| false | false | false  |

- `!`

| x     | !x    |
|-------|-------|
| true  | false |
| false | true  |

#### Identità booleane
- Idenpotenza: `a && a = a`
- Doppia negazione: `!!a = a`
- Assorbimento: `a && (a||b) = a`

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
### Esercizi
- Scrivere un programma che dato il prezzo di un bene e data l'aliquota IVA, determina l'imponibile (`calcolo_imponibile.go`)

## Controllo del flusso
### if/else if/else
```go 
if <>{
    <corpo>
}else if <>{
    <corpo>
}else{
    <corpo>
}
```
#### Esercizi
- Leggi 2 frazioni positive e stabilisci se la prima è minore della seconda (`controllo_frazione_minore.go`)
- Date 2 date stabilire se la prima precede la seconda (`controllo_date.go`)
- Dati a, b, c stabilire se l'equazione ax^2+bx+c = 0 ha 2 soluzioni reali distinte (`ax^2+bx+c.go`)
- Stabilire se la somma di 2 interi ha la cifra delle decine uguale a 7 (`controllo_decine_7.go`)
### for
- for 1-ario (unario)
```go
for A{
    <corpo>
}
```
#### Esercizi
- (`MCD_poveri.go`)
- (`MCD.go`)
- (`somma_n_naturali.go`)
- (`gauss_iterato.go`)
- for 0-ario
```go
for {
    <corpo>
}
```
- for 3-ario
```go
for A; B; C{
    <corpo>
}
```
#### Esercizi
- Dato n, stampa i primi n quadrati perfetti (`n_quadratti_perfetti.go`)
- Dato n, stampa le prime n potenze di 2 (`n_potenze_2.go`)

Stampa di pattern
- Dato n, stampa un quadrato di * con lato n (`quadrato_lato_n.go`)
- Dato n, stampa un triangolo reattangolo con i cateti lunghi n (`tri_ret_cateti_n.go`)

Lettura ripetuta di input
- Date n persone, calcola l'altezza media delle persone (`h_media_n_persone.go`)

Lettura con "tappo"
- calcolare la media di tutti i numeri inseriti fino all'inserimento della cifra 0 (`calcolo_media_tappo.go`)
- Stbilire se un numero è primo (`numero_primo_tappo.go`)

### 🛑Break
Interrompe l'esecuzione del ciclo in cui è contenuto
```go
for <>{
    if <>{
        break
    }
}
```
- Emersione rapida
> Con emersione rapida (early return o fail-fast) si riferisce a una tecnica di programmazione in cui le condizioni di errore o i casi speciali vengono gestiti immediatamente, permettendo al flusso principale del codice di rimanere pulito e focalizzato sulla logica principale.
#### Esercizi
- 

### 🏃Continue
Forza la prossima esecuzione del ciclo
```go
for <>{
    if <>{
        continue
    }
}
```
#### Esercizi
- Dato n, stampa i numeri primi <n (`numeri_primi_minori_n.go`)
- Stampa i primi n numeri primi (`)

## 1️⃣ / 0️⃣ Bool
### Operatori booleani
- AND: `&&`

| x     | y     | x && y |
|-------|-------|--------|
| true  | true  | true   |
| true  | false | false  |
| false | true  | false  |
| false | false | false  |

- OR: `||`

| x     | y     | x \|\| y |
|-------|-------|--------|
| true  | true  | true   |
| true  | false | true   |
| false | true  | true   |
| false | false | false  |

- NOT: `!`

| x     | !x    |
|-------|-------|
| true  | false |
| false | true  |

> Negli operatori di bool la precedenza segue quest'ordine: ! > && > ||
### Identità booleane
- Idenpotenza: `a && a = a`
- Doppia negazione: `!!a = a`
- Assorbimento: `a && (a||b) = a`
### Leggi di De Morgan
- `!(x && y) = !x || !y`
- `!(x || y) = !x && !y`
### Short-circuit evaluation
> Lo short-circuit evaluation è una tecnica utilizzata dai linguaggi di programmazione, inclusi Go, per ottimizzare l'esecuzione di espressioni logiche.
> Fa riferimento al comportamento in cui l'esecuzione di un'operazione logica si interrompe non appena il risultato finale può essere determinato senza dover valutare ulteriori condizioni.

## Esercizi
- Stabilire se un numero termina con 3 zeri
- Basandosi sull'anno di nascita e sull'anno corrente, stabilire se una persona è maggiorenne
- Stabilire se la somma delle cifre è <10
- Dati 2 numeri stampali in ordine crescente
- Dati 3 numeri stampali in ordine crescente
- Stabilire se un anno è bisestile

## Funzioni
