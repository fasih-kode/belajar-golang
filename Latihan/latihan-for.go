package main

import "fmt"

func main() {
	// 1. Perulangan dasar
	//
	//Buat program dengan for yang mencetak angka 1 sampai 10.
	fmt.Println("=== Soal 1 ===")
	for i := 0; i < 10; i++ {
		fmt.Println("Ini angka :", i+1)
	}
	fmt.Println()

	// 2. Penjumlahan
	//
	//Gunakan for untuk menghitung jumlah semua angka dari 1 sampai 100 lalu cetak hasilnya.
	fmt.Println("=== Soal 2 ===")
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println()

	// 3. Kelipatan
	//
	//Gunakan for untuk mencetak semua angka kelipatan 3 dari 1 sampai 30.

	// Cetak di dalam loop → kamu melihat proses bertahap.
	//
	//Cetak di luar loop → kamu melihat hasil akhir saja.
	fmt.Println("=== Soal 3 ===")
	for i := 3; i <= 30; i++ {
		i += 3
		fmt.Println(i)
	}
	fmt.Println()

	// 4. Mirip while
	//
	//Gunakan for (tanpa init & post) untuk menggandakan sebuah angka mulai dari 2, sampai nilainya lebih besar dari 100. Cetak hasil akhirnya.
	fmt.Println("=== Soal 4 ===")
	a := 2
	for a < 100 {
		a += a
		fmt.Println(a)
	}
	fmt.Println()

	// 5. Loop string
	//
	//Gunakan for range untuk mencetak setiap huruf dari string "golang" beserta index-nya.
	fmt.Println("=== Soal 5 ===")
	var nama = "Golang"
	for i, v := range nama {
		fmt.Println("Indeks", i, "Value", string(v))
	}
	fmt.Println()

	// 6. Loop slice
	//
	//Buat slice []int{5, 10, 15, 20}.
	//Gunakan for range untuk mencetak jumlah semua elemen di slice tersebut.
	fmt.Println("=== Soal 6 ===")
	nums := []int{5, 10, 15, 20}
	tam := 0
	for _, v := range nums {
		tam += v
	}
	fmt.Println(tam)
	fmt.Println()

	// 7. Infinite loop
	//
	//Buat program dengan for {} yang terus berjalan dan mencetak "Hello".
	//(Pakai break supaya bisa berhenti setelah 5 kali).
	fmt.Println("=== Soal 7 ===")
	count := 0
	for {
		fmt.Println("Hello")
		count++

		if count == 5 {
			break
		}
	}
	// Penjelasan:
	//
	//count := 0 → variabel penghitung.
	//
	//for {} → loop tanpa batas.
	//
	//count++ → setiap putaran, count bertambah 1.
	//
	//if count == 5 { break } → kalau sudah 5 kali, keluar dari loop.

}
