package main

import (
	"fmt"
)

// Struct untuk menyimpan informasi klub
type Klub struct {
	nama string
}

// Struct untuk menyimpan hasil pertandingan
type Pertandingan struct {
	skorA int
	skorB int
}

// Fungsi utama
func main() {
	var klubA, klubB Klub
	var hasil []string
	var pertandingan []Pertandingan

	// Input nama klub
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA.nama)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB.nama)

	// Input skor hingga skor negatif
	i := 1
	for {
		var p Pertandingan
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&p.skorA, &p.skorB)

		// Cek jika ada skor negatif → berhenti
		if p.skorA < 0 || p.skorB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		// Simpan hasil pertandingan ke slice
		pertandingan = append(pertandingan, p)

		// Cek hasil pertandingan
		if p.skorA > p.skorB {
			hasil = append(hasil, klubA.nama)
		} else if p.skorB > p.skorA {
			hasil = append(hasil, klubB.nama)
		} else {
			hasil = append(hasil, "Draw")
		}

		i++
	}

	// Tampilkan hasil rekap
	for j, h := range hasil {
		fmt.Printf("Hasil %d : %s\n", j+1, h)
	}
}