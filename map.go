package main

import "fmt"

func main() {

	//membuat map
	mahasiswa := map[string]int{
		"Ani":  90,
		"Budi": 85,
		"Cici": 78,
	}
	//Akses nilai dengan key
	fmt.Println("Nilai Ani:", mahasiswa["Ani"])

	// Tambah data baru
	mahasiswa["Doni"] = 90

	//Ubah nilai
	mahasiswa["Budi"] = 95

	// Hapus data
	delete(mahasiswa, "Cici")

	// Iterasi dengan for range
	for nama, nilai := range mahasiswa {
		fmt.Println(nama, ":", nilai)
	}
}
