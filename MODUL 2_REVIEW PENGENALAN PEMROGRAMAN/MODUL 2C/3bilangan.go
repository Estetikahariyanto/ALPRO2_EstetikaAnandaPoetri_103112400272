package main

import (
	"fmt"
)

// Fungsi untuk mendapatkan faktor dari sebuah bilangan
func getFactors(n int) []int {
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// Fungsi untuk mengecek apakah bilangan prima
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

	// Cek apakah bilangan prima
	if isPrime(b) {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}