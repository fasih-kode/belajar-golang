package main

import "fmt"

func main() {

	var usia int
	fmt.Println("Masukkan usia kamu: ")
	fmt.Scanln(&usia)

	if usia < 0 || usia > 120 {
		fmt.Println("Data tidak valid")
	} else if usia >= 0 && usia <= 12 {
		fmt.Println("Tingkat anak-anak")
	} else if usia >= 13 && usia <= 17 {
		fmt.Println("Tingkat remaja")
	} else if usia >= 18 && usia <= 59 {
		fmt.Println("Tingkat dewasa")
	} else {
		fmt.Println("Tingkat lansia")
	}
}
