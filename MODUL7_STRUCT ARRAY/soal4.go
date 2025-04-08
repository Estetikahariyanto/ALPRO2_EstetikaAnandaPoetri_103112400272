package main

import (
	"fmt"
)

// Konstanta maksimum panjang array
const NMAX int = 127

// Struct Tabel, menyimpan array karakter dan panjang isi array
type Tabel struct {
	tab [NMAX]rune
	n   int
}

// Fungsi input isi array
func isiArray(t *Tabel) {
	fmt.Print("Masukkan teks (tanpa spasi): ")
	var input string
	fmt.Scanln(&input)

	t.n = len(input)
	for i, ch := range input {
		t.tab[i] = ch
	}
}

// Fungsi mencetak isi array
func cetakArray(t Tabel) {
	for i := 0; i < t.n; i++ {
		fmt.Printf("%c ", t.tab[i])
	}
	fmt.Println()
}

// Fungsi membalik isi array
func balikArray(t *Tabel) Tabel {
	var hasil Tabel
	hasil.n = t.n
	for i := 0; i < t.n; i++ {
		hasil.tab[i] = t.tab[t.n-1-i]
	}
	return hasil
}

// Fungsi untuk mengecek apakah array adalah palindrom
func palindrome(t Tabel) bool {
	for i := 0; i < t.n/2; i++ {
		if t.tab[i] != t.tab[t.n-1-i] {
			return false
		}
	}
	return true
}

// Fungsi utama
func main() {
	var t Tabel

	// Input array karakter
	isiArray(&t)

	// Cetak array asli
	fmt.Print("Teks asli     : ")
	cetakArray(t)

	// Balik isi array dan cetak
	tBalik := balikArray(&t)
	fmt.Print("Teks reverse  : ")
	cetakArray(tBalik)

	// Cek apakah palindrom
	if palindrome(t) {
		fmt.Println("Palindrome    : true")
	} else {
		fmt.Println("Palindrome    : false")
	}
}