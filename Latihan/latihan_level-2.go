package main

import (
	"fmt"
	"strconv"
)

func main() {
	daftarBarang := []map[string]string{}

	fmt.Println("===== Menu Toko =====")
	fmt.Println("1. Input Barang")
	fmt.Println("2. Daftar Barang")
	fmt.Println("3. Cari Barang")
	fmt.Println("4. Hapus Barang")
	fmt.Println("5. Update Barang")
	fmt.Println("6. Total Nilai Inventaris")
	fmt.Println("7. Keluar")

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
		case "4":
			var hapus string
			fmt.Print("Masukkan barang yang akan dihapus: ")
			fmt.Scan(&hapus)

			ditemukan := false
			indeksHapus := -1

			// cari indeks barang yang akan dihapus
			for i, barang := range daftarBarang {
				if barang["nama"] == hapus {
					indeksHapus = i
					ditemukan = true
					break
				}
			}
			if ditemukan {
				// hapus barang dari slice
				daftarBarang = append(daftarBarang[:indeksHapus], daftarBarang[indeksHapus+1:]...)
				fmt.Println("Barang", hapus, "Berhasil dihapus")
			} else {
				fmt.Println("Barang", hapus, "tidak ditemukan, tidak ada yang dihapus")
			}
		case "5":
			var update string
			fmt.Print("Masukkan nama barang yang akan diupdate: ")
			fmt.Scan(&update)

			ditemukan := false
			indeksUpdate := -1

			// cari indeks barang yang akan diupdate
			for i, barang := range daftarBarang {
				if barang["nama"] == update {
					indeksUpdate = i
					ditemukan = true
					break
				}
			}

			if ditemukan {
				var namaUpdate, hargaUpdate, stokUpdate string
				fmt.Print("Masukkan nama barang baru: ")
				fmt.Scan(&namaUpdate)
				fmt.Print("Masukkan harga barang baru: ")
				fmt.Scan(&hargaUpdate)
				fmt.Print("Masukkan stok barang baru: ")
				fmt.Scan(&stokUpdate)

				// update data barang
				daftarBarang[indeksUpdate]["nama"] = namaUpdate
				daftarBarang[indeksUpdate]["harga"] = hargaUpdate
				daftarBarang[indeksUpdate]["stok"] = stokUpdate

				fmt.Println("Barang", update, "Berhasil diupdate")
			} else {
				fmt.Println("Barang", update, "tidak ditemukan, tidak ada yang diupdate")
			}
		case "6":
			fmt.Println("\n===== Total Nilai Inventaris =====")
			if len(daftarBarang) == 0 {
				fmt.Println("Barang masih kosong")
			} else {
				totalBarang := len(daftarBarang)
				totalNilai := 0

				for _, barang := range daftarBarang {
					harga, errHarga := strconv.Atoi(barang["harga"])
					stok, errStok := strconv.Atoi(barang["stok"])

					if errHarga == nil && errStok == nil {
						totalNilai += harga * stok
					}
				}

				fmt.Println("Total Jumlah Barang:", totalBarang, "jenis")
				fmt.Println("Total Nilai Inventaris: Rp", totalNilai)
				fmt.Println("==================================")
			}
		case "7":
			fmt.Println("Keluar dari program...")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih 1-7.")
		}
	}
}
