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