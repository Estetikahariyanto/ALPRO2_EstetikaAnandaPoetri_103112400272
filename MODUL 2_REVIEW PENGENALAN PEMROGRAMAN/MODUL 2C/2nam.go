package main
import "fmt"
func main() {
	var nam float64
	var nmk string
	fmty.Print("nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)
	if nam > 80 {
		nam = "A"
	}
	if nam > 72.5 {
		nam = "AB"
	}
if nam > 65 {
	nam = "B"
}
if nam > 57.5 {
	nam = "BC"
}
if nam > 50 {
	nam = "C"
}
if nam > 40 {
	nam = "D"
} else if nam <= 40 {
	nam = "E"
}
fmt.Println("nilai mata kuliah: ", nmk)
}