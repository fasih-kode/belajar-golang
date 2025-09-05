package main

import "fmt"

// cara 1: pakai var
var nama string = "Fasih"

// cara 2: tanpa tipe (Go otomatis deteksi)
var umur = 41
var kata = "durian"
var des = 2.7

// Gunakan := → kalau lagi di dalam function, ini yang paling umum.
//
//Gunakan var ... = → kalau di luar function, atau kalau Anda mau biar jelas bahwa itu var, bukan konstanta.
//
//Gunakan var ... type = → kalau tipenya penting ditulis eksplisit (misalnya int64 untuk database, atau float32 untuk grafis), biar tidak salah.

var lokasi = "Pekandangan"

// buat variabel langsung banyak

var (
	namadepan    = "Fasih"
	namabelakang = "Khuddin"
)

func main() {
	kota := "Indramayu"
	//cara 3: singkat (hanya di dalam funciton dan var bentuk ini harus dalam main)
	jumlah := 24
	persen := 2.5
	fmt.Println(nama, umur, kota, jumlah, persen, kata, des, lokasi, namadepan, namabelakang)
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	// fungsi string, menghitung jumlah karakter
	fmt.Println(len("Fasikhuddin"))

	// mengambil karakter pada posisi yang ditentukan
	fmt.Println("Fasikhuddin"[0])
	fmt.Println("Fasikhuddin"[3])
}
