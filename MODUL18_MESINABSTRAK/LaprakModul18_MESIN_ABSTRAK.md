# <h1 align="center">Laporan Praktikum Modul 18 <br>MESIN ABSTRAK</h1>
<p align="center">ESTETIKA ANANDA POETRI HARIYANTO - 103112400272</p>

## Dasar Teori
Mesin abstrak adalah model komputasi yang dirancang di atas model mesin komputasi yang telah ada, yaitu tipe data dan operasi-operasi dasarnya dibuat menggunakan tipe data dan operasi-operasi yang tersedia di mesin di bawahnya. Teknik ini merupakan salah satu cara untuk membangun perangkat lunak.

## NO 1
![[soal.png]]

```go
package main
import (
"fmt"
"strings"
)
var input string 
var idx int       
var currentChar byte   
func start() {
idx = 0
currentChar = input[idx]
}
func maju() {
idx++
if idx < len(input) {
currentChar = input[idx]
}
}
func eop() bool {
    return currentChar == '.'
}
func getcc() byte {
return currentChar
}
func main() { 
fmt.Print("Masukkan string (akhiri dengan titik): ")
fmt.Scanln(&input)
 if !strings.HasSuffix(input, ".") {
 fmt.Println("Input harus diakhiri dengan titik ('.').")
 return
}
start()
totalChars := 0
countA := 0
countLE := 0
var prevChar byte = 0
for !eop() {
ch := getcc()
totalChars++
if ch == 'A' || ch == 'a' {
countA++        
if prevChar == 'L' && ch == 'E' {
 countLE++
}
prevChar = ch
maju()
}
fmt.Println("\n=== Hasil pembacaan ===")
fmt.Println("Jumlah karakter terbaca:", totalChars)
fmt.Println("Jumlah huruf 'A':", countA)
if totalChars > 0 {
fmt.Printf("Frekuensi huruf 'A': %.2f%%\n", float64(countA)/float64(totalChars)*100)
}
fmt.Println("Jumlah kata 'LE':", countLE)
}
```

>Output![[MODUL18_MESINABSTRAK/output 1/ss1.png]]
  
Program ini membaca string diakhiri dengan tanda titik (.) lalu menggunakan mesin karakter utk membaca satu satu karakter, yg kedua adalah menghitung statistik yaitu jumlah karakter, jumlah huruf A, frekuensi huruf A, dan jumlah pasangan (LE).
`currentChar`: menyimpan karakter yang sedang dibaca.
`input`: string yang akan dibaca, diinput dari pengguna.
`idx`: posisi saat ini dari karakter yang sedang dibaca.







