package main

import "fmt"

func main() {

	var nama string
	var usia int
	var nilai int

	fmt.Print("Masukkan nama: ", nama)
	fmt.Scanln(&nama)
	fmt.Print("Masukkan usia: ", usia)
	fmt.Scanln(&usia)
	fmt.Print("Masukkan nilai: ", nilai)
	fmt.Scanln(&nilai)

	if nilai < 0 || nilai > 100 {
		println(" Nilai tidak valid")
	} else if nilai >= 90 {
		println(" Nilai A")
	} else if nilai >= 80 {
		println(" Nilai B")
	} else if nilai >= 70 {
		println(" Nilai C")
	} else {
		println(" Nilai D")
	}

	var kategori string
	switch {
	case usia >= 0 && usia <= 12:
		kategori = "anak-anak"
	case usia >= 13 && usia <= 17:
		kategori = "Remaja"
	case usia >= 18 && usia <= 59:
		kategori = "Dewasa"
	default:
		kategori = "Lansia"
	}

	fmt.Printf("Nama: %s\n", nama)
	fmt.Printf("Usia: %d (%s)\n", usia, kategori)
	fmt.Printf("Nilai: %d\n", nilai)

	mapel := []string{"Matematika", "Bahasa", "IPA"}
	for _, v := range mapel {
		fmt.Println(v)
	}
}
