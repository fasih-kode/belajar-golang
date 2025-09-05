package main

import "fmt"

func main() {
	// cara 1: pakai var
	var nama string = "Fasih"

	// cara 2: tanpa tipe (Go otomatis deteksi)
	var umur = 41
	var kata = "durian"
	var des = 2.7

	//cara 3: singkat (hanya di dalam funciton)
	kota := "Indramayu"
	jumlah := 24
	persen := 2.5

	fmt.Println(nama, umur, kota, jumlah, persen, kata, des)
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	// fungsi string, menghitung jumlah karakter
	fmt.Println(len("Fasikhuddin"))

	// mengambil karakter pada posisi yang ditentukan
	fmt.Println("Fasikhuddin"[0])
	fmt.Println("Fasikhuddin"[3])
}
