package main

import "fmt"

func main() {
	daftarBarang := []map[string]string{}

	fmt.Println("===== Menu Toko =====")
	fmt.Println("1. Input Barang")
	fmt.Println("2. Daftar Barang")
	fmt.Println("3. Cari Barang")
	fmt.Println("4. Hapus Barang")
	fmt.Println("5. Keluar")

	for {
		var input string
		fmt.Print("Pilih Nomor: ")
		fmt.Scan(&input)
		switch input {
		case "1":
			var n int
			fmt.Print("Jumlah Barang: ")
			fmt.Scan(&n)
			for i := 0; i < n; i++ {
				var nama, harga, stok string
				fmt.Print("Masukkan nama barang: ")
				fmt.Scan(&nama)
				fmt.Print("Masukkan harga barang: ")
				fmt.Scan(&harga)
				fmt.Print("Masukkan stok barang: ")
				fmt.Scan(&stok)

				// buat map baru
				barangBaru := map[string]string{
					"nama":  nama,
					"harga": harga,
					"stok":  stok,
				}
				daftarBarang = append(daftarBarang, barangBaru)
				fmt.Println("Barang Berhasil Ditambahkan")
			}
		case "2":
			fmt.Println("===== Daftar Barang =====")
			if len(daftarBarang) == 0 {
				fmt.Println("Barang masih kosong")
			} else {
				for i, val := range daftarBarang {
					fmt.Println("Barang yang ke-", i+1)
					fmt.Println("Nama : ", val["nama"])
					fmt.Println("Harga :", val["harga"])
					fmt.Println("Stok :", val["stok"])
					fmt.Println("=============================")
				}
			}
		case "3":
			var cari string
			fmt.Print("\nMasukkan nama barang yang dicari: ")
			fmt.Scan(&cari)

			ditemukan := false
			for _, barang := range daftarBarang {
				if barang["nama"] == cari {
					fmt.Println("Ditemukan!")
					fmt.Println("Nama :", barang["nama"])
					fmt.Println("Harga :", barang["harga"])
					fmt.Println("Stok :", barang["stok"])
					ditemukan = true
					break
				}
			}

			if !ditemukan {
				fmt.Println("Maaf, barang", cari, "tidak ada di data")
			}
		}

	}

}
