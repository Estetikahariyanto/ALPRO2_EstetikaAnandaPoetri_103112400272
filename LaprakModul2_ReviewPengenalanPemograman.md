# <h1 align="center">Laporan Praktikum Modul 2 <br> REVIEW PENGENALAN PEMROGRAMAN</h1>
<p align="center">ESTETIKA ANANDA POETRI HARIYANTO - 103112400272</p>

## Dasar Teori

 Golang yaitu bahasa pemrograman yang dirancang untuk efisiensi dan kesederhanaan. Struktur dasarnya mencakup package, import, dan fungsi main(). Input diperoleh dengan fmt.Scanln(), dan pertukaran nilai variabel menggunakan variabel untuk rotasi data.

## Unguided

## SOAL 2A
### NO 1

Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?

```go
package main

import "fmt"

func main() {
    var (
        satu, dua, tiga string
        temp            string
    )
    fmt.Print("Masukkan input string pertama: ")
    fmt.Scanln(&satu)
    fmt.Print("Masukkan input string kedua: ")
    fmt.Scanln(&dua)
    fmt.Print("Masukkan input string ketiga: ")
    fmt.Scanln(&tiga)
    fmt.Println("Output awal =", satu, dua, tiga)
    temp = satu
    satu = dua
    dua = tiga
    tiga = temp
    fmt.Println("Output akhir =", satu, dua, tiga)

}
```
> Output
![Output](./Output/gambar1.png)



Program ini menerima tiga string dari pengguna, Melakukan rotasi nilai secara searah jarum jam dan Menampilkan hasil sebelum dan sesudah rotasi.

### NO 2

```go
package main

import (
    "fmt"
)
func isLeapYear(year int) bool {
    if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
        return true
    }
    return false
}
func main() {
    var year int
    fmt.Print("Tahun: ")
    fmt.Scan(&year)
    fmt.Printf("Kabisat: %t\n", isLeapYear(year))
}
```

> Output
![Output](./Output/gambar2.png)

Program ini menentukan apakah suatu tahun adalah **kabisat** atau bukan berdasarkan aturan pembagian bilangan.
Digunakan **operator modulus (`%`)** untuk memeriksa habis atau tidaknya suatu tahun dibagi angka tertentu.
Hasil yang ditampilkan adalah `true` untuk kabisat dan `false` untuk bukan kabisat.

### NO 3

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    var r float64
    fmt.Print("Jejari = ")
    fmt.Scan(&r)
    volume := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
    luas := 4 * math.Pi * math.Pow(r, 2)
    fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
```

> Output
![Output](./Output/gambar3.png)

Program ini membaca **jari jari bola**, lalu menghitung **volume** dan **luas permukaan bola** menggunakan rumus matematika yang sesuai.
Fungsi `math.Pow()` digunakan untuk melakukan operasi pangkat.
Hasil ditampilkan dengan format 4 angka desimal untuk lebih presisi.

### NO 4

```go
package main

import (
    "fmt"
)
func main() {
    var celsius float64

    //input user
    fmt.Print("Temperatur Celsius: ")
    fmt.Scan(&celsius)
    //suhu
    reamur := celsius * 4 / 5
    fahrenheit := (celsius * 9 / 5) + 32
    kelvin := celsius + 273
    //output
    fmt.Printf("Derajat Reamur: %.0f\n", reamur)
    fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
    fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
```

> Output
![Output](./Output/gambar4.png)

Membaca input suhu dalam Celsius menggunakan `fmt.Scanln(&celsius)`.
Menggunakan rumus konversi untuk menghitung suhu dalam Reamur, Fahrenheit, dan Kelvin.
Menampilkan hasil konversi menggunakan `fmt.Printf()` utk hasil memiliki 2 angka di belakang koma untuk presisi.

### NO 5.

```go
package main

import (

    "fmt"

)
func main() {
    var asciiCodes [5]int
    var chars [3]rune
    fmt.Println("Masukkan 5 angka ASCII:")
    for i := 0; i < 5; i++ {
        fmt.Scan(&asciiCodes[i])
    }
    fmt.Println("Masukkan 3 karakter tanpa spasi:")
    for i := 0; i < 3; i++ {
        fmt.Scanf("%c", &chars[i])
    }
    fmt.Print("Hasil ASCII ke karakter: ")
    for i := 0; i < 5; i++ {
        fmt.Printf("%c", asciiCodes[i])
    }
    fmt.Println()
    fmt.Print("Hasil karakter: ")
    for i := 0; i < 3; i++ {
        fmt.Printf("%c", chars[i])
    }
    fmt.Println()
}
```


> Output
![Output](./Output/gambar5.png)

Program ini menunjukkan gimana data numerik (bilangan integer) itu dapat direpresentasikan sebagai karakter pakai kode ASCII.


## SOAL 2B

## NO 1

```go
package main
import (
    "fmt"
)
func main() {
    //urutan warna
    expected := [4]string{"merah", "kuning", "hijau", "ungu"}
    for i := 1; i <= 5; i++ {
        var input [4]string
        fmt.Printf("Percobaan %d: ", i)
        //input 4 warna
        for j := 0; j < 4; j++ {
            fmt.Scan(&input[j])
        }
        match := true
        for j := 0; j < 4; j++ {
            if input[j] != expected[j] {
                match = false
                break
            }
        }
        //tampilkan
        if match {
            fmt.Println("true")
        } else {
            fmt.Println("false")
        }
    }
}
```

> Output
![Output](./Output/gambar6.png)

Program ini bertujuan untuk mengevaluasi keberhasilan percobaan kimia berdasarkan urutan warna cairan dalam tabung reaksi. Jika urutan warna dalam 5 percobaan berturut-turut adalah merah, kuning, hijau, ungu, maka hasilnya BERHASIL. Jika ada satu percobaan yang tidak sesuai, maka hasilnya FALSE.

### NO 2

```go
package main

import (
    "fmt"
    "strings"
)
func main() {

    var n int
    fmt.Print("Masukkan jumlah bunga (N): ")
    fmt.Scan(&n)
    var bungaList []string
    for i := 1; i <= n; i++ {
        var bunga string
        fmt.Printf("Bunga %d: ", i)
        fmt.Scan(&bunga)
        if strings.ToUpper(bunga) == "SELESAI" {
            break
        }
        bungaList = append(bungaList, bunga)
    }
    pita := strings.Join(bungaList, " - ")
    fmt.Println("Pita:", pita)
    fmt.Println("Jumlah Bunga:", len(bungaList))
}
```

> Output
![Output](./Output/gambar7.png)

Program ini menerima input jumlah bunga N lalu meminta N nama bunga utk disusun dalam pita memakai tanda "-" sebagai pemisah. jika N = 0, maka pita kosong.
Modifikasi program memungkinkan input berhenti saat pengguna mengetik SELESAI, lalu menampilkan isi pita dan jumlah bunga yang dimasukkan.

### NO 3

```go
package main

import (
    "fmt"
)

func main() {
    var kantong1, kantong2 float64
    for {
        fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
        fmt.Scan(&kantong1, &kantong2)
        if kantong1 >= 9 || kantong2 >= 9 {
            fmt.Println("Salah satu kantong sudah mencapai 9 kg atau lebih. Proses selesai.")
            break
        }
    }
}
```

> Output
![Output](./Output/gambar8.png)

Program ini meminta pengguna untuk memasukkan berat dua kantong terpal secara berulang hingga salah satu kantong memiliki berat minimal 9 kg, lalu program akan berhenti dengan pesan bahwa proses selesai.

### NO 4

```go
package main

import (
    "fmt"
    "math"
)

func f(k int) float64 {
    numerator := math.Pow(float64(4*k+2), 2)
    denominator := float64((4*k+1)*(4*k+3))
    return numerator / denominator
}

func approximateSqrt2(K int) float64 {
    result := 1.0
    for k := 0; k < K; k++ {
        result *= f(k)
    }
    return result
}
func main() {
    var K int
    fmt.Print("Masukkan nilai K: ")
    fmt.Scan(&K)
    fk := f(K)
    fmt.Printf("Nilai f(K) = %.10f\n", fk)
    approxSqrt2 := approximateSqrt2(K)
    fmt.Printf("Nilai akar 2 = %.10f\n", approxSqrt2)
}
```

> Output
![Output](./Output/gambar9.png)

Program ini menghitung f(K) pakai rumus matematika, lalu menggunakannya untuk memperkirakan √2 dengan mengalikan hasil f(k) dari k = 0 sampai K-1. Semakin besar K, semakin mendekati nilai √2.

## SOAL 2C

### NO 1

```go
package main
import (
    "fmt"
)
func hitungBiaya(berat int) int {
    biayaPerKg := 10000
    biayaTambahanRingan := 5
    biayaTambahanBerat := 15
    totalKg := berat / 1000
    sisaGram := berat % 1000
    biaya := totalKg * biayaPerKg
    if sisaGram > 0 {
        if sisaGram <= 500 {
            biaya += sisaGram * biayaTambahanRingan
        } else {
            biaya += sisaGram * biayaTambahanBerat
        }
    }
    return biaya
}
func main() {
    var berat int
    fmt.Print("Masukkan berat paket (gram): ")
    fmt.Scan(&berat)
    biaya := hitungBiaya(berat)
    fmt.Printf("Total biaya pengiriman: Rp. %d\n", biaya)
}
```

> Output
![Output](./Output/gambar10.png)

Program ini tujuannya untuk menghitung biaya pengiriman berdasarkan berat paket yang di kasihkan dri pengguna. masukan berat paket dlm satuan gram dri pengguna. kemudian program akan mengkonversi berat tsb ke satuan kilogram dan menghitung sisa berat dlm gram, lalu.. biaya dasar pengiriman ditetapkan sebesar 10.000/kilo. jika ada sisa berat sampai 500 gram maka biaya tambahan dihitung Rp.5/gram kalau sisa berat lebih dari 500 gram, biaya tambahannya yaitu Rp. 15/gram. jika total berat paket lebih dari 10 kg program akan membulatkannya ke lipatan 10 kg utk perhitungan biaya.
dannn program menampilkan total biaya pengiriman yg hrus dibayar.

### NO 2

```go
package main

import "fmt"
func main() {
    var nam float64
    var nmk string
    fmty.Print("nilai akhir mata kuliah: ")
    fmt.Scanln(&nam)
    if nam > 80 {
        nam = "A"
    }
    if nam > 72.5 {
        nam = "AB"
    }
if nam > 65 {
    nam = "B"
}
if nam > 57.5 {
    nam = "BC"
}
if nam > 50 {
    nam = "C"
}
if nam > 40 {
    nam = "D"
} else if na <= 40 {
    nam = "E"
}
fmt.Println("nilai mata kuliah: ", nmk)

}
```

##### A. Jika NAM diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?

NAM = 80.1, program akan mengevaluasi kondisi `if nam > 80 { nmk = "A" }`, jadi nilai NMK yang dihasilkan adalah "A".

Eksekusi sesuai, karna 80.1 lebih besar dari 80, sehingga masuk ke kategori A berdasarkan tabel nilai.

##### B. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!

Penggunaan `if` yang terpisah tanpa `else if`
> karena smua kondisi pakai if, nah program tetap meriksa semua pernyataan meskipun salah satu kondisi sudah terpenuhi.
> program menggunakan else if utk memastikan hanya satu kondisi yg dievaluasi.

Kesalahan dlm batas rentang
> kondisi if nam > 72.5 {nmk = "AB"} tidak mencakup nilai di atas 80, tetapi di program asli, tidak ada batas atas yg eksplisit.
> jadii jika NAM = 90, program bisa menetapkan "A" , tapi msih di evaluasi if nam > 72.5 dan menetapkan "AB" yg salah.

Rentang angka tidak tertutup dengan baik
> utk nilai NAM = 50, tdk ada kondisi else if nam <= 57.5 && nam > 50, sehingga program akan melewatkan kondisi untuk nilai C.

Perbaikan alur program nya :
> pakai else if utk memastikan hanya satu kondisi yang dievaluasi.
>memastikan rentang nilai terdefinisi dengan jelas yg sesuai.

##### C. Perbaiki program tersebut! Ujilah dengan masukan: 93.5, 70.6, dan 49.5. Seharusnya keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’.

Program yang sudah di perbaiki :
```go
package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah:", nmk)
}
```

### NO 3

```go
package main

import (
    "fmt"
)

func getFactors(n int) []int {
    var factors []int
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            factors = append(factors, i)
        }
    }
    return factors
}
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
func main() {
    var b int
    fmt.Print("Masukkan bilangan bulat > 1: ")
    fmt.Scan(&b)
    if b <= 1 {
        fmt.Println("Bilangan harus lebih dari 1")
        return
    }
    factors := getFactors(b)
    fmt.Printf("Bilangan: %d\n", b)
    fmt.Print("Faktor: ")
    for _, factor := range factors {
        fmt.Printf("%d ", factor)
    }
    fmt.Println()
    if isPrime(b) {
        fmt.Println("Prima: true")
    } else {
        fmt.Println("Prima: false")
    }
}
```

> Output
![Output](./Output/gambar11.png)

Untuk menerima input sebuah bilangan bulat lebih dari 1, mencari faktor-faktornya, dan menentukan apakah bilangan tersebut merupakan bilangan prima atau bukan..

