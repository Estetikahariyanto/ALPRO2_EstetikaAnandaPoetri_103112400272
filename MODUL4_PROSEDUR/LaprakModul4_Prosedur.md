# <h1 align="center">Laporan Praktikum Modul 4 <br> PROSEDUR</h1>
<p align="center">ESTETIKA ANANDA POETRI HARIYANTO - 103112400272</p>

## Dasar Teori

 Prosedur adalah subprogram yang terdiri dari kumpulan instruksi yang dapat dipanggil dalam program utama untuk mengurangi kompleksitas kode. Tidak seperti fungsi, prosedur tidak memiliki nilai kembalian dan tidak menggunakan kata kunci Return.
## Unguided

## NO 1

![[MODUL4_PROSEDUR/gambarsoal/soal1.png]]

```go
package main
import (
    "fmt"
)
func factorial(n int) int {
    hasil := 1
    for i := 1; i <= n; i++ {
        hasil *= i
    }
    return hasil
}
func permutation(n, r int) int {
    if n < r {
        return 0
    }
    return factorial(n) / factorial(n-r)
}
func combination(n, r int) int {
    if n < r {
        return 0
    }
    return factorial(n) / (factorial(r) * factorial(n-r))
}
func main() {
    var a, b, c, d int
    fmt.Print("Masukkan empat bilangan (a b c d): ")
    fmt.Scan(&a, &b, &c, &d)
    Pac := permutation(a, c)
    Cac := combination(a, c)
    Pbd := permutation(b, d)
    Cbd := combination(b, d)
    fmt.Println(Pac, Cac)
    fmt.Println(Pbd, Cbd)
}
```

>Output
> ![Output](MODUL4_PROSEDUR/Output/ss1.png)

Program ini untuk menghitung nilai permutasi (P(n,r)) dan kombinasi (C(n,r)) dari dua pasang bilangan yang diberikan sebagai input.
Untuk menjalankan program ini, digunakan tiga prosedur utama, yaitu:

1. Prosedur faktorial yang menghitung nilai faktorial dari suatu bilangan.
2. Prosedur permutasi yang menggunakan hasil faktorial untuk menghitung permutasi suatu bilangan.
3. Prosedur kombinasi yang juga menggunakan faktorial untuk menghitung kombinasi.
### NO 2

![[MODUL4_PROSEDUR/gambarsoal/soal2.png]]

```go
package main

import (
    "fmt"
    "strings"
)
type Peserta struct {
    Nama    string
    Soal    int
    Waktu   int
}
func hitungSkor(input []string) Peserta {
    var pemenang Peserta
    for _, line := range input {
        if line == "Selesai" {
            break
        }]
        fields := strings.Fields(line)
        if len(fields) < 2 {
            continue
        }
        nama := fields[0]
        waktuTotal := 0
        jumlahSoal := 0
        for _, waktuStr := range fields[1:] {
            var waktu int
            fmt.Sscanf(waktuStr, "%d", &waktu)
            if waktuTotal+waktu > 301 {
                break
            }
            waktuTotal += waktu
            jumlahSoal++
        }

        if jumlahSoal > pemenang.Soal || (jumlahSoal == pemenang.Soal && waktuTotal < pemenang.Waktu) {
            pemenang = Peserta{Nama: nama, Soal: jumlahSoal, Waktu: waktuTotal}
        }
    }
    return pemenang
}
func main() {
    input := []string{
        "Astuti 20 50 301 61 71 75 10",
        "Bertha 25 47 36 61 66 60 65 21",
        "Selesai",
    }
    pemenang := hitungSkor(input)
    fmt.Printf("%s %d %d\n", pemenang.Nama, pemenang.Soal, pemenang.Waktu)

}
```

> Output
> ![Output](MODUL4_PROSEDUR/Output/ss2.png)

Program untuk menentukan pemenang dalam kompetisi pemrograman berdasarkan jumlah soal yang berhasil diselesaikan dalam waktu maksimal 301 menit. Jika ada lebih dari satu peserta dengan jumlah soal yang sama, pemenang ditentukan berdasarkan waktu penyelesaian yang lebih sedikit.

### NO 3

Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah ½n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir.

![[soal3.png]]

```go
package main

import (
    "fmt"
func cetakDeret(n int) {
    for n != 1 {
        fmt.Printf("%d ", n)
        if n%2 == 0 {
            n /= 2
        } else {
            n = 3*n + 1
        }
    }

    fmt.Println(n)

}
func main() {
    var n int
    fmt.Print("Masukkan bilangan positif (< 1000000): ")
    fmt.Scan(&n)
    if n <= 0 || n >= 1000000 {
        fmt.Println("Bilangan harus lebih dari 0 dan kurang dari 1000000.")
        return
    }
    cetakDeret(n)

}
```

> Output
>  ![Output](MODUL4_PROSEDUR/Output/ss3.png)

Program yang diminta bertujuan untuk mencetak deret bilangan berdasarkan aturan yang diberikan oleh Skiena dan Revilla dalam _Programming Challenges_. Deret ini dimulai dengan sebuah bilangan bulat positif _n_ dan mengikuti aturan berikut:

1. Jika _n_ adalah bilangan genap, maka suku berikutnya adalah _n/2_.
2. Jika _n_ adalah bilangan ganjil, maka suku berikutnya adalah _3n + 1_.
3. Proses ini terus berlanjut hingga bilangan tersebut mencapai angka 1, di mana deret akan berhenti.



