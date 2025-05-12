# <h1 align="center">Laporan Praktikum Modul 17 <br>SKEMA PEMROSESAN SEKUENSIAL</h1>
<p align="center">ESTETIKA ANANDA POETRI HARIYANTO - 103112400272</p>

## Dasar Teori
1. skema pemrosesan sekuensial
yaitu pendekatan pemrograman dimana data diproses satu per satu secara berurutan. dgn bantuan struktur perulangan atau looping dan percabangan skema ini dpt menyelesaikannya.

2. pembacaan data tanpa marker
semua data yang masuk pasti diproses, jumlah data biasanya diketahui terlebih dahulu, dan cocok digunakan saat jumlah data sudah diketahui sblmnya.

3. pembacaan data dengan masker
4. kasus rangkaian data kosong
5. elemen pertama diproses tersendiri

## Guided

## NO 1
Aldi memiliki daftar nilai ulangan matematika temannya: 75, 60, 90, 85, dan 70. Ia ingin mengurutkan nilai tersebut dari yang terkecil ke yang terbesar menggunakan **metode Bubble Sort**.
Pertanyaan:
1. Tunjukkan proses pengurutan nilai menggunakan Bubble Sort hingga semua nilai terurut.
2. Berapa kali pertukaran (swap) terjadi dalam proses ini?

```go
package main
import (
    "fmt"
)
func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
        fmt.Printf("setelah pass %d: %v\n", i+1, arr)
    }
}

func main() {
    nilai := []int{75, 60, 90, 85, 70}
    fmt.Println("nilai sebelum diurutkan:", nilai)
    bubbleSort(nilai)
    fmt.Println("nilai setelah diurutkan:", nilai)
}
```

>Output![[ssguided.png]]
  
fungsi bubbleSort(arr []int) ini untuk menerima sebuah slice arr berisi angka dan n menyimpan jumlah elemen dalam slice.

Program ini untuk mengurutkan nilai-nilai ulangan matematika dari yang terkecil ke yang terbesar menggunakan metode Bubble Sort.

1. Pertama, program menyimpan daftar nilai ulangan dalam sebuah variabel `nilai`, yaitu: `75, 60, 90, 85, dan 70`.
2. Kemudian, daftar nilai ini dikirim ke fungsi `bubbleSort`, yang bertugas mengurutkan data tersebut.
3. Di dalam fungsi `bubbleSort`, program melakukan perbandingan antara dua angka yang bersebelahan. Jika angka di sebelah kiri lebih besar dari yang di kanan, maka keduanya ditukar.
4. Proses perbandingan dan pertukaran ini diulang beberapa kali (disebut pass) sampai seluruh data benar-benar dalam urutan yang tepat.
5. Setelah setiap pass (putaran pengurutan), program menampilkan kondisi daftar nilai saat itu, agar kita bisa melihat prosesnya langkah demi langkah.
6. Setelah semua data selesai diurutkan, program menampilkan hasil akhirnya, yaitu daftar nilai yang sudah tersusun dari yang terkecil ke yang terbesar.

#### 2 Berapa kali pertukaran (swap) terjadi dalam proses ini?
[75, 60, 90, 85, 70]
proses bubble sort dan jumlah pertukaran
ke 1:
- 75 > 60 > tukar > [60,75,90,85,70] 1 swap
- 75 < 90 > tidak ditukar
- 90 > 85 > tukar > [60,75,85,90,70] 2 swap
- 90 > 70 > tukar > [60,75,85,70,90] 3 swap
ke 2:
- 60 < 75 > tidak tukar
- 75 < 85 > tidak tukar
- 85 > 70 > tukar > [60,75,70,85,90] 4 swap
ke 3:
- 60 < 75 > tidak tukar
- 75 > 70 > tukar > [60,70,75,85,90] 5 swap
ke 4:
- 60 < 70 > tidak tukar
- 70 < 75 > tidak tukar
total SWAP = 
3+1+1+0 = 5 kali

## Unguided

## NO 1

Diberikan sejumlah bilangan real yang diakhiri dengan marker 9999, cari rerata dari bilangan bilangan tsb.

```go
package main

import (
    "fmt"
)

func main() {
    var number float64
    var sum float64
    var count int
    fmt.Println("masukkan bilangan real satu per satu (akhiri dengan 9999):")

    for {
        fmt.Print("-> ")
        fmt.Scan(&number)
        if number == 9999 {
            break
        }
        sum += number
        count++
    }
    if count > 0 {
        average := sum / float64(count)
        fmt.Printf("Rerata dari bilangan yang dimasukkan adalah: %.2f\n", average)
    } else {
        fmt.Println("Tidak ada bilangan yang dimasukkan.")
    }
}
```

>Output![[MODUL17_SKEMA_PEMROSESAN_SEKUENSIAL/output/ss1.png]]

Program ini menerima bilangan satu per satu dari pengguna, menghitung jumlah dan banyaknya bilangan, dan ketika 9999 dimasukkan..program akan menghitung rata rata dan menampilkannya.


## NO 2

Diberikan string x dan n buah string. x adalah data pertama yang dibaca, n adalah data bilangan yang dibaca kedua, dan n data berikutnya adalah data string. buat algoritma untuk menjawab pertanyaan berikut:
a. apakah string x ada dalam kumpulan n data string tsb ?
b. pada posisi ke brp string x tsb ditemukan ?
c. ada berapakah string x dalam kumpulan n data string tsb ?
d. adakah sedikitnya dua string x dalam n data string tsb ?

```go
package main

import (
    "fmt"
)
func main() {
    var x string
    var n int
    fmt.Print("masukkan string x yang dicari: ")
    fmt.Scan(&x)
    fmt.Print("masukkan jumlah data string n: ")
    fmt.Scan(&n)
    data := make([]string, n)
    fmt.Println("masukkan", n, "string:")
    for i := 0; i < n; i++ {
        fmt.Printf("data ke-%d: ", i+1)
        fmt.Scan(&data[i])
    }
    ditemukan := false
    var posisi []int
    for i, val := range data {
        if val == x {
            ditemukan = true
            posisi = append(posisi, i+1) 
        }
    }
    if ditemukan {
        fmt.Println("string ditemukan.")
        fmt.Print("ditemukan pada posisi: ")
        for _, pos := range posisi {
            fmt.Print(pos, " ")
        }
        fmt.Println()
    } else {
        fmt.Println("string tidak ditemukan dalam data.")
    }
}
```

>Output![[MODUL17_SKEMA_PEMROSESAN_SEKUENSIAL/output/ss2.png]]

A. apakah string x ada dalam kumpulan n data string tsb ?
jawaban nya : ya, string mangga ada.

B. pada posisi ke brp string x tsb ditemukan ?
jawaban nya : ke 2 & 5

C. ada berapakah string x dalam kumpulan n data string tsb ?
jawaban nya : ada 2 string yg bernama mangga dlm kumpulan data tsb.

D. adakah sedikitnya dua string x dalam n data string tsb ?
jawaban nya : ada, krena mangga muncul sebanyak 2 kali.

## NO 3

![[MODUL17_SKEMA_PEMROSESAN_SEKUENSIAL/gambarsoal/soal1.png]]

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)
func main() {
    var totalDrops int
    const mmPerDrop = 0.0001
    fmt.Print("Masukkan jumlah tetesan air hujan: ")
    fmt.Scan(&totalDrops)
    rand.Seed(time.Now().UnixNano())
    var countA, countB, countC, countD int
    for i := 0; i < totalDrops; i++ {
        x := rand.Float64()
        y := rand.Float64()
        if x < 0.5 && y < 0.5 {
            countA++
        } else if x >= 0.5 && y < 0.5 {
            countB++
        } else if x < 0.5 && y >= 0.5 {
            countC++
        } else {
            countD++
        }
    }
    rainA := float64(countA) * mmPerDrop
    rainB := float64(countB) * mmPerDrop
    rainC := float64(countC) * mmPerDrop
    rainD := float64(countD) * mmPerDrop
    fmt.Printf("Curah hujan daerah A: %.4f millimeter\n", rainA)
    fmt.Printf("Curah hujan daerah B: %.4f millimeter\n", rainB)
    fmt.Printf("Curah hujan daerah C: %.4f millimeter\n", rainC)
    fmt.Printf("Curah hujan daerah D: %.4f millimeter\n", rainD)
}
```

>Output![[MODUL17_SKEMA_PEMROSESAN_SEKUENSIAL/output/ss3.png]]

Program ini bertujuan untuk logika dan pemodelan simulasi sederhana dalam konteks yang nyata, yaitu mengukur curah hujan berdasarkan sebaran tetesan air hujan.







