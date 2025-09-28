package main

import "fmt"

func main() {

	daftarBarang := []map[string]string{}

	for {
		fmt.Println("========= Menu Toko Sederhana =========")
		fmt.Println("1. Tambah Barang Baru")
		fmt.Println("2. Lihat Daftar Barang")
		fmt.Println("3. Cari Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Keluar")

		var input string
		fmt.Println("Pilih Nomor: ")
		fmt.Scan(&input)
		switch input {
		case "1":
			var jumlah int
			fmt.Print("Jumlah: ")
			fmt.Scan(&jumlah)

			for i := 0; i < jumlah; i++ {
				var nama, harga, stok string
				fmt.Print("Nama barang: ")
				fmt.Scan(&nama)
				fmt.Print("Harga barang: ")
				fmt.Scan(&harga)
				fmt.Print("stok: ")
				fmt.Scan(&stok)

				barangBaru := map[string]string{
					"nama":  nama,
					"harga": harga,
					"stok":  stok,
				}
				daftarBarang = append(daftarBarang, barangBaru)
				fmt.Println("Barang berhasil ditambahkan", daftarBarang)
			}
		case "2":
			if len(daftarBarang) == 0 {
				fmt.Println("Belum ada barang")
			} else {
			fmt.Println("===== Daftar Barang =====")
			for i, v := range daftarBarang {
				fmt.Printf("Barang ke-[%v] = %v\n", i, v)
				fmt.Println("Nama: ", v["nama"])
				fmt.Println("Harga: ", v["harga"])
				fmt.Println("Stok: ", v["stok"])
				fmt.Println("=======================")
			}
		}
	}
}