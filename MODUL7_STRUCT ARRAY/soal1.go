package main

import (
	"fmt"
	"math"
)

// Struct Titik
type Titik struct {
	x int
	y int
}

// Struct Lingkaran
type Lingkaran struct {
	pusat Titik
	r     int
}

// Fungsi menghitung jarak antar dua titik
func jarak(p1, p2 Titik) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
}

// Fungsi mengecek apakah titik di dalam lingkaran
func dalamLingkaran(l Lingkaran, p Titik) bool {
	return jarak(l.pusat, p) <= float64(l.r)
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titik Titik

	// Input lingkaran 1
	fmt.Scan(&lingkaran1.pusat.x, &lingkaran1.pusat.y, &lingkaran1.r)

	// Input lingkaran 2
	fmt.Scan(&lingkaran2.pusat.x, &lingkaran2.pusat.y, &lingkaran2.r)

	// Input titik yang dicek
	fmt.Scan(&titik.x, &titik.y)

	// Cek posisi titik
	dalam1 := dalamLingkaran(lingkaran1, titik)
	dalam2 := dalamLingkaran(lingkaran2, titik)

	// Output berdasarkan hasil
	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}