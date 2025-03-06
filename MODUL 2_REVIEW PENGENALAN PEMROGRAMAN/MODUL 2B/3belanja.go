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