package main

import (
	"fmt"
	"math"
)

// Struct untuk menyimpan array bilangan bulat
type DataArray struct {
	data []int
}

// Menampilkan seluruh isi array
func (d DataArray) TampilkanSemua() {
	fmt.Println("Isi array:", d.data)
}

// Menampilkan elemen dengan indeks ganjil
func (d DataArray) TampilkanIndeksGanjil() {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < len(d.data); i += 2 {
		fmt.Print(d.data[i], " ")
	}
	fmt.Println()
}

// Menampilkan elemen dengan indeks genap
func (d DataArray) TampilkanIndeksGenap() {
	fmt.Print("Indeks genap: ")
	for i := 0; i < len(d.data); i += 2 {
		fmt.Print(d.data[i], " ")
	}
	fmt.Println()
}

// Menampilkan elemen dengan indeks kelipatan x
func (d DataArray) TampilkanIndeksKelipatan(x int) {
	fmt.Printf("Indeks kelipatan %d: ", x)
	for i := 0; i < len(d.data); i++ {
		if i%x == 0 {
			fmt.Print(d.data[i], " ")
		}
	}
	fmt.Println()
}

// Menghapus elemen pada indeks tertentu
func (d *DataArray) HapusIndeks(i int) {
	if i < 0 || i >= len(d.data) {
		fmt.Println("Indeks tidak valid!")
		return
	}
	d.data = append(d.data[:i], d.data[i+1:]...)
	fmt.Println("Setelah dihapus:")
	d.TampilkanSemua()
}

// Menghitung rata-rata
func (d DataArray) RataRata() float64 {
	total := 0
	for _, val := range d.data {
		total += val
	}
	return float64(total) / float64(len(d.data))
}

// Menghitung standar deviasi
func (d DataArray) StandarDeviasi() float64 {
	rata := d.RataRata()
	var total float64
	for _, val := range d.data {
		total += math.Pow(float64(val)-rata, 2)
	}
	return math.Sqrt(total / float64(len(d.data)))
}

// Menghitung frekuensi suatu nilai
func (d DataArray) Frekuensi(nilai int) int {
	count := 0
	for _, val := range d.data {
		if val == nilai {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := DataArray{data: make([]int, n)}

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr.data[i])
	}

	arr.TampilkanSemua()
	arr.TampilkanIndeksGanjil()
	arr.TampilkanIndeksGenap()

	// Tampilkan indeks kelipatan x
	var x int
	fmt.Print("Masukkan bilangan untuk kelipatan indeks: ")
	fmt.Scan(&x)
	arr.TampilkanIndeksKelipatan(x)

	// Hapus indeks tertentu
	var idx int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)
	arr.HapusIndeks(idx)

	// Rata-rata
	fmt.Printf("Rata-rata: %.2f\n", arr.RataRata())

	// Standar deviasi
	fmt.Printf("Standar deviasi: %.2f\n", arr.StandarDeviasi())

	// Frekuensi nilai tertentu
	var cari int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)
	fmt.Printf("Frekuensi %d: %d kali\n", cari, arr.Frekuensi(cari))
}