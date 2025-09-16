package main

import "fmt"

func sapa(nama string) {
	fmt.Println("Halo!", nama)
}

func hitung(angka int) {
	fmt.Println("Ini adalah angka", angka)
}

func tambah(a, b int) int {
	return a + b
}

func kali(a, b int) int {
	return a * b
}

func bagi(a, b float32) float32 {
	return float32(a / b)
}

func main() {
	sapa("Fasih")

	hitung(5)

	hasilTambah := tambah(70, 23)
	fmt.Println("Hasil tambah adalah:", hasilTambah)

	hasilKali := kali(70, 23)
	fmt.Println("Hasil kali adalah:", hasilKali)

	hasilBagi := bagi(70, 23)
	fmt.Println("Hasil bagi adalah:", hasilBagi)

}
