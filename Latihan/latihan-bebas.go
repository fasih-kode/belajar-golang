package main

import "fmt"

func main() {
	var jumlah int
	fmt.Print("Enter jumlah: ")
	fmt.Scanln(&jumlah)

	// Buat Slice untuk barang
	barang := make([]string, 0)

	// Buat map untuk nama dan harga barang
	harga := make(map[string]int)

	// Buat perulangan input
	for i := 0; i < jumlah; i++ {
		var name string
		var hrg int
		fmt.Print("Enter name: ")
		fmt.Scanln(&name)
		fmt.Print("Enter harga: ")
		fmt.Scanln(&hrg)
		// Masukkan name ke dalam slice barang
		barang = append(barang, name)
		// Masukkan hrg ke dalam map
		harga[name] = hrg
	}
	// daftar barang yang dibeli
	fmt.Println("\n=== Daftar Belanja ===")
	total := 0
	for _, name := range barang {
		fmt.Printf("%s Rp.%d\n", name, harga[name])
		total += harga[name]
	}
	// Total belanja
	fmt.Printf("\n=== Total Belanja ===")
	fmt.Printf("\nTotal Rp.%d\n", total)
}
