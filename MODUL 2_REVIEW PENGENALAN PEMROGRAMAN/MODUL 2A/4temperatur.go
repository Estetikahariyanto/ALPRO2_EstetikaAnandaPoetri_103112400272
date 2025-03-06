package main
import (
	"fmt"
)
func main() {
	var celsius float64

	//input user
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)

	//suhu
	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273

	//output
	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}