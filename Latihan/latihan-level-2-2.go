package main

import "fmt"

func main() {

	// masukkan jumlah barang yang dicatat
	var jumlah int
	fmt.Print("Masukkan jumlah barang: ")
	fmt.Scanln(&jumlah)

	// Slice untuk menampung nama barang
	barang := make([]string, 0)

	// Map untuk menampung barang dan harganya
	harga := make(map[string]int)

	// input data barang
	for i := 0; i < (jumlah); i++ {
		var nama string
		var hrg int
		fmt.Print("Nama barang: ")
		fmt.Scanln(&nama)
		fmt.Print("Harga: ")
		fmt.Scanln(&hrg)

		barang = append(barang, nama) // simpan ke slice
		harga[nama] = hrg             // simpan ke map
	}

	// Tampilkan daftar belanja
	fmt.Println("\n=== Daftar Belanja ===")
	total := 0
	for _, nama := range barang {
		fmt.Printf("%s: Rp.%d\n", nama, harga[nama])
		total += harga[nama]
	}

	// Tampilkan total
	fmt.Printf("\nTotal Belanja: Rp.%d\n", total)

	// Percabangan untuk kategori belanja
	if total > 100000 {
		fmt.Println("Belanja banyak sekali!")
	} else if total > 50000 {
		fmt.Println("Belanja lumayan")
	} else {
		fmt.Println("Belanja hemat")
	}

	// Switch berdasarkan jumlah barang
	switch jumlah {
	case 0:
		fmt.Println("Keranjang kosong")
	case 1:
		fmt.Println("Belanja hanya satu barang")
	case 2, 3:
		fmt.Println("Belanja beberapa barang")
	default:
		fmt.Println("Belanja banyak barang")
	}

}
