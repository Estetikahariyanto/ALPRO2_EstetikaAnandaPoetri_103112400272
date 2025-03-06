package main
import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

func main() {
	var year int
	fmt.Print("Tahun: ")
	fmt.Scan(&year)

	fmt.Printf("Kabisat: %t\n", isLeapYear(year))
}