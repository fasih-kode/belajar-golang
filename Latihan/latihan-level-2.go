package main

import "fmt"

func main() {

	// Jumlah data siswa
	var jumlah int
	fmt.Println("Masukkan jumlah data siswa: ")
	fmt.Scanln(&jumlah)

	// Map untuk menyimpan nama dan nilai
	siswa := make(map[string]int)

	// input data siswa
	for i := 0; i < jumlah; i++ {
		var nama string
		var nilai int
		fmt.Println("Nama Siswa: ")
		fmt.Scanln(&nama)
		fmt.Println("Nilai Siswa: ")
		fmt.Scanln(&nilai)

		siswa[nama] = nilai
	}

	// Jumlahkan semua data siswa
	fmt.Println("\n=== Data Nilai Siswa === ")
	total := 0
	for nama, nilai := range siswa {
		fmt.Printf("%s: %d\n", nama, nilai)
		total += nilai

		// Menentuka grade pakai if
		var grade string
		if nilai >= 90 {
			grade = "A"
		} else if nilai >= 75 {
			grade = "B"
		} else if nilai >= 60 {
			grade = "C"
		} else {
			grade = "D"
		}

		// komentar pakai switc
		switch grade {
		case "A":
			fmt.Println("Sangat Baik")
		case "B":
			fmt.Println("Bagus, Pertahankan")
		case "C":
			fmt.Println("Cukup, Perlu belajar lagi")
		case "D":
			fmt.Println("Kurang, harus lebih giat")
		}
	}

	// hitung rata-rata
	rata := float64(total) / float64(jumlah)
	fmt.Printf("\nRata-rata nilai: %.2f\n", rata)
}
