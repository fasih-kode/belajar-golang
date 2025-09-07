package main

import "fmt"

func main() {
	var angka int

	fmt.Println("Masukkan angka kamu: ")
	fmt.Scan(&angka)

	switch angka {
	case 1:
		fmt.Println("Kamu masuk Senin")
	case 2:
		fmt.Println("Kamu masuk Selasa")
	case 3:
		fmt.Println("kamu masuk Rabu")
	case 4:
		fmt.Println("Kamu masuk Kamis")
	case 5:
		fmt.Println("kamu masuk Jumat")
	case 6:
		fmt.Println("Kamu masuk Sabtu")
	case 7:
		fmt.Println("Kamu masuk Ahad")
	default:
		fmt.Println("Data tidak valid")
	}
}
