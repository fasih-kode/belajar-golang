package main

import "fmt"

func main() {

	var angka int
	fmt.Println("Masukkan angka anda: ")
	fmt.Scanln(&angka)

	if angka < 0 {
		fmt.Println("Angka tidak valid")
	} else if angka%2 == 0 {
		fmt.Println("Angka genap")
	} else {
		fmt.Println("Angka ganjil")
	}
}
