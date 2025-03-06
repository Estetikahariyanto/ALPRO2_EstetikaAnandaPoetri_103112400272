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