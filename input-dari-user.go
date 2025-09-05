package main

import "fmt"

func main() {
	var nama string
	var umur int
	var alamat string
	fmt.Println("masukkan nama kamu: ")
	fmt.Scanln(&nama)
	fmt.Println("masukkan umur kamu: ")
	fmt.Scanln(&umur)
	fmt.Println("masukkan alamat: ")
	fmt.Scanln(&alamat)

	fmt.Printf("Halo %s, umur kamu %d, alamat %s\n", nama, umur, alamat)
}
