package main

import "fmt"

func main() {
	var nama, harga, stok string

	var daftarBarang []map[string]string

	fmt.Print("Masukkan nama barang: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan harga barang: ")
	fmt.Scan(&harga)
	fmt.Print("Masukkan stok: ")
	fmt.Scan(&stok)

	barangBaru := make(map[string]string{
		"nama" : nama,
		"harga" : harga,
		"stok" : stok,
	}

	daftarBarang = append(daftarBarang, barangBaru)

	fmt.Print("Daftar Barang ", daftarBarang)
}
