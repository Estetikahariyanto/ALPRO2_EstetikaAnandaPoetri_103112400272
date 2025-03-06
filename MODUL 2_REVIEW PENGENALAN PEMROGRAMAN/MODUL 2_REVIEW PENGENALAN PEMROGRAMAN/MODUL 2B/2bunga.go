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