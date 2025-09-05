package main

import "fmt"

func main() {
	// cara 1: pakai var
	var nama string = "Fasih"

	// cara 2: tanpa tipe (Go otomatis deteksi)
	var umur = 41

	//cara 3: singkat (hanya di dalam funciton)
	kota := "Indramayu"

	fmt.Println(nama, umur, kota)
}
