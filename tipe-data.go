package main

import "fmt"

func main() {
	// 1. deklarasi variabel
	var umur int = 25
	var nama string = "Fasih"
	var aktif bool = true
	var pi float64 = 3.14159265359
	var reang string = "kata"

	// cetak tipe dan nilai
	fmt.Printf("Type: %T Value: %v\n", umur, umur)
	fmt.Printf("Type: %T Value: %v\n", nama, nama)
	fmt.Printf("Type: %T Value: %v\n", aktif, aktif)
	fmt.Printf("Type: %T Value: %v\n", pi, pi)

	// 2. operasi sederhana
	fmt.Println("Umur + 5 =", umur+5)
	fmt.Println("Pi * 2 =", pi*2)
	fmt.Println("Nama lengkap =", nama+" Reang")
	fmt.Println("Apakah umur > 20 ?", umur > 20)

	// 3. konversi
	var angka float64 = float64(umur) / 3.0
	fmt.Println("Umur dibagi 3 =", angka)
	fmt.Println("Panjang nama =", len(nama))
	fmt.Println(reang)
}
