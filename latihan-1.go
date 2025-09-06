package main

import "fmt"

// Fungsi untuk luas persegi panjang
// var p dan l disini hanya sementara, menunggu input dari user
func hitungLuasPersegiPanjang(p, l float64) float64 {
	return p * l
}

// Fungsi untuk keliling persegi panjang
func hitungKelilingPersegiPanjang(p, l float64) float64 {
	return 2 * (p + l)

}

func main() {
	// buat variabel panjang dan lebar
	var panjang, lebar float64

	// user diminta masukkan panjang dan lebar
	fmt.Println("Masukkan Panjang Persegi Panjang: ")
	fmt.Scanln(&panjang)
	fmt.Println("Masukkan Lebar Persegi Panjang: ")
	fmt.Scanln(&lebar)

	// buat var luas dan keliling
	luas := hitungLuasPersegiPanjang(panjang, lebar)
	keliling := hitungKelilingPersegiPanjang(panjang, lebar)

	// menampilkan hasil perhitungan
	fmt.Printf("Luas persegi panjang dengan panjang %.2f dan lebar %.2f adalah %.2f\n", panjang, lebar, luas)
	fmt.Printf("Keliling Persegi Panjang dengan panjang %.2f dan lebar %.2f adalah %.2f\n", panjang, lebar, keliling)
}
