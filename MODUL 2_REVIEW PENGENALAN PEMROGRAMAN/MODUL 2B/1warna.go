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