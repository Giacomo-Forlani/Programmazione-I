# Corso di Programmazione I dell'UniMi

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
### Tipi di Basi
- Interi
    - Con segno
        - int
        - int8
        - int16
        - int32
        - int64
    - Senza segno
        - uint
        - uint8
        - uint16
        - uint32
        - uint64
- Floating Point
    - Reali
        - float32
        - float64
    - Complessi
        - complex64
        - complex128
- Altri
    - string
    - bool
    - byte
    - rune
    - uintptr
    - error

#### Interi
- Notazione decimale
- Notazione esadecimale
- notazione ottale
##### Interi con segno
- Rappresentazione in complemento a 2

###### Esercizi
- Dato n trova qual è il prossimo numero della sequenza di Collatz (`n_successivo_seq_collatz.go`)
- Date 2 date trova la differenza esatta in giorni tra le 2 (`differenza_date.go`)
- Dato n stampa il triangolo vuoto di base e altezza n (`triangolo_vuoto.go`)

##### Interi senza segno
- Valori: interi ≥ 0
- Rappresentazione posizionale binaria: 
- Overflow

###### Esercizi
- Stampa la potenza di 3 (`potenza_3.go`)

#### Floating point
- Notazione scientifica

#### Rune
- US-ASCII (7bit)
- Unicode
    - Unicode encoding scheme: rune -> UTF-32 (Uint32)
    - UTF-8: rappresentazione di discusione variabile
    - Carattere: sequenza di 1, 2, 3 o 4 byte
    - 1° byte: numero di byte che seguono

### Dichiarazioni di variabile
- Dichiarazione di variabile singola: `var <VARIABILE> <TIPO>`
- Dichiarazione di variabile multipla `var <VARIABILE>,<VARIABILE>,<VARIABILE> <TIPO>`

### Shadowing
> Le variabili dichiarate esistono solo all'interno della funzione in cui sono state dichiarate e non in quelle esterne o in quelle contenute

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
- Maggiore(>): `>`
- Maggiore Uguale (≥): `>=`
- Minore(<): `<`
- Minore Uguale (≤): `<=`
- Uguale (=): `==`
- Diverso (≠): `!=`

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

```go
fmt.Scan(&a)
```
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
#### for 1-ario (unario)
```go
for A{
    <corpo>
}
```
##### Esercizi
- (`MCD_poveri.go`)
- (`MCD.go`)
- (`somma_n_naturali.go`)
- (`gauss_iterato.go`)

#### for 0-ario
```go
for {
    <corpo>
}
```
#### for 3-ario
```go
for A; B; C{
    <corpo>
}
```
##### Esercizi
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

### 🛑 Break
Interrompe l'esecuzione del ciclo in cui è contenuto
```go
for <>{
    if <>{
        break
    }
}
```
> Emersione rapida (early return o fail-fast): si riferisce a una tecnica di programmazione in cui le condizioni di errore o i casi speciali vengono gestiti immediatamente, permettendo al flusso principale del codice di rimanere pulito e focalizzato sulla logica principale.

#### Esercizi
- 

### 🏃 Continue
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
- Stampa i primi n numeri primi (`n_numeri_primi.go`)

### Switch
```go
switch <selettore>{
case <espr>, <espr>, ...:
    <codice>
case <espr>, <espr>, ...:
    <codice>
default:
    <codice>
}
```
> l'opzione default è opzionale

<esercizi>Esercizi<esercizi>
- Dato un numero tra 0 e 100, nonn compresi, scrivi una fumzione che restituisce il numero in lettere (`n_in_lettere.go`)
- Dato un numero, stampa un carta di un mazzo di da poker (`carte_poker.go`)
- Dato un numero, stampa un carta di un mazzo di da poker utilizzando la libreria strconv (`alt_carte_poker.go`)

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
```go
func <nome funzione> (<parametri formali>) <tipi di ritorno>{
    <corpo della funzione>
}
```

### Esercizi
- Dato n intero, scrivi una funziona che restituisca il quadrato di n (`func_quadrato_n.go`)
- Scrivi una fuzione che calcola l'ipotenusa dati 2 cateti, utilizzando la libreria math (`func_calc_ipotenusa.go`)
- Dato n intero, scrivi una funzione che determina se n è pari (`func_n_pari.go`)
- Scrivere una funzione che determina se un numero è divisore di un altro (`func_divisore.go`)
- Scrivere una funzione che calcola quanti zeri compaiono nella scrittura di un numero. Usala per calcolare quanti 0 ci sono se scrivete tutti i numeri da 1 a 1.000.000 (`func_conta_0.go`)
- Scrivete una funzione che, date le coordinate di due punti sul piano cartesiano, restituisca il coefficiente angolare della retta passante (`func_coef_ang.go`)
- Dato n, determina se è un numero primo di Mersen. Usala per scrivere i primi n numeri di Mersen (`n_primi_mersen.go`)

### Funzioni che restituiscono più valori

<style>
    esercizi{
        color: red;
    }
    h1{
        color: purple;
    }
</style>
