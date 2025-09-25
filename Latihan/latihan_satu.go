package main

import "fmt"

// Fungsi untuk hitung diskon, pakai named return
func hitungDiskon(total int) (diskon int) {
	if total > 100000 {
		diskon = total / 10
	}
	return
}

func main() {
	for {
		var input string
		fmt.Print("\nKetik 'exit' untuk keluar, atau 'lanjut' untuk belanja: ")
		fmt.Scan(&input)

		if input == "exit" {
			fmt.Println("bye")
			break
		}

		var jumlah int
		fmt.Print("Enter jumlah item: ")
		fmt.Scan(&jumlah)

		// variabel total harga
		total := 0

		// perulangan input barang
		for i := 0; i < jumlah; i++ {
			var namaBarang string
			var harga, qty int

			fmt.Printf("item ke %d\n ", i)
			fmt.Print("nama barang: ")
			fmt.Scan(&namaBarang)
			fmt.Print("harga: ")
			fmt.Scan(&harga)
			fmt.Print("Banyak: ")
			fmt.Scan(&qty)

			subtotal := harga * qty
			total += subtotal
		}

		// hitung diskon
		discount := hitungDiskon(total)

		// hasil akhir
		fmt.Println("\n===== STRUK BELANJA =====")
		fmt.Println("Total Belanja: ", total)
		fmt.Println("Discount: ", discount)
		fmt.Println("Harus Dibayar :", total-discount)
		fmt.Println("============================")
	}
}
