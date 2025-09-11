package main

import "fmt"

func main() {

	buah := []string{"Apel", "Mangga", "Jeruk"}
	fmt.Println("Daftar buah:", buah)

	//Menambah data baru
	buah = append(buah, "Pisang")
	fmt.Println("setelah ditambah:", buah)

	//Mengakses elemen
	fmt.Println("Buah pertama:", buah[0])
	fmt.Println("Buah terakhit:", buah[len(buah)-1])

	keranjang := []string{}
	keranjang = append(keranjang, "sabun", "sikat", "odol")
	fmt.Println(len(keranjang))
	fmt.Println(keranjang)
	fmt.Println(keranjang[len(keranjang)-1])

	// hapus index 1 → "sikat"
	keranjang = append(keranjang[:1], keranjang[2:]...)
	// artinya ambil sabun dan odol
	fmt.Println("Setelah hapus index 1:", keranjang)

	// reset isi keranjang
	keranjang = []string{"sabun", "sikat", "odol"}

	// hapus index 0 → "sabun"
	keranjang = append(keranjang[:0], keranjang[1:]...)
	// artinya keranjang satu kosong dan keranjang dua sikat odol
	fmt.Println("Setelah hapus index 0:", keranjang)

	// reset isi keranjang
	keranjang = []string{"sabun", "sikat", "odol"}

	// hapus index 2 → "odol"
	keranjang = append(keranjang[:2], keranjang[3:]...)
	// artinya ambil sabun sikat, dan keranjang kosong
	fmt.Println("Setelah hapus index 2:", keranjang)

	hari := []string{}
	hari = []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Ahad"}

	// 2 hari pertama
	fmt.Println("Tampilkan hanya 2 hari pertama:", hari[:2])

	// 3 hari terakhir
	fmt.Println("Tampilkan hanya 3 hari terakhir:", hari[4:])
}
