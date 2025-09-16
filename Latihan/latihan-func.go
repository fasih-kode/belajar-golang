// 1. Struktur file .go
package main

// 2. Import package
import "fmt"

// 3. Komentar
// Program Kasir Sederhana
// Latihan gabungan Level 1–3 Golang

// 4. Variabel global & 5. Tipe data dasar
var namaKasir string // string
var totalBelanja int // int
var diskon float64   // float
var isMember bool    // bool

// 6. Konstanta
const PPN = 0.11 // 11%

// 16. Fungsi
// Fungsi hitungTotal (menjumlahkan harga barang)
func hitungTotal(harga []int) int {
	total := 0
	for _, h := range harga {
		total += h
	}
	return total
}

// Fungsi hitungDiskon (menghitung diskon)
func hitungDiskon(total int) float64 {
	if total >= 500000 {
		return 0.2 // diskon 20%
	} else if total >= 200000 {
		return 0.1 // diskon 10%
	} else if total >= 100000 {
		return 0.05 // diskon 5%
	} else {
		return 0.0 // tidak ada diskon
	}
}

// Fungsi bagi (multiple return untuk cicilan)
func bagi(total int, cicilan int) (int, error) {
	if cicilan == 0 {
		return 0, fmt.Errorf("jumlah cicilan tidak boleh 0")
	}
	return total / cicilan, nil
}

func main() {
	// 7. Input dari user → nama pembeli
	var namaPembeli string
	fmt.Print("Masukkan nama pembeli: ")
	fmt.Scanln(&namaPembeli)

	// 8. Input jumlah barang
	var jumlahBarang int
	fmt.Print("Masukkan jumlah barang: ")
	fmt.Scanln(&jumlahBarang)

	// 12. Array kasir
	kasir := [3]string{"Andi", "Budi", "Citra"}
	namaKasir = kasir[0] // default kasir pertama

	// 13. Slice untuk nama barang
	barang := make([]string, 0)

	// 14 & 15. Map untuk barang + harga
	harga := make(map[string]int)

	// Slice harga untuk hitung total
	hargaList := make([]int, 0)

	// 10. for → input barang & harga
	for i := 0; i < jumlahBarang; i++ {
		var namaBarang string
		var harganya int
		fmt.Printf("Masukkan nama barang ke-%d: ", i+1)
		fmt.Scanln(&namaBarang)
		fmt.Printf("Masukkan harga %s: ", namaBarang)
		fmt.Scanln(&harganya)

		barang = append(barang, namaBarang)
		harga[namaBarang] = harganya
		hargaList = append(hargaList, harganya)
	}

	// 11. for range → tampilkan daftar belanja
	fmt.Println("\n=== Daftar Belanja ===")
	for _, b := range barang {
		fmt.Printf("%s : Rp.%d\n", b, harga[b])
	}

	// Hitung total belanja (pakai fungsi)
	totalBelanja = hitungTotal(hargaList)
	fmt.Println("Total Belanja:", totalBelanja)

	// 9. if else if else → tentukan diskon
	diskon = hitungDiskon(totalBelanja)
	potongan := int(float64(totalBelanja) * diskon)
	fmt.Println("Diskon:", potongan)

	// Tambahkan PPN
	ppn := int(float64(totalBelanja) * PPN)
	totalAkhir := totalBelanja - potongan + ppn
	fmt.Println("PPN:", ppn)
	fmt.Println("Total Bayar (setelah diskon + PPN):", totalAkhir)

	// 10. switch → metode pembayaran
	var metode int
	fmt.Println("\nPilih metode pembayaran:")
	fmt.Println("1. Cash")
	fmt.Println("2. Transfer")
	fmt.Println("3. E-Wallet")
	fmt.Print("Pilihan Anda: ")
	fmt.Scanln(&metode)

	switch metode {
	case 1:
		fmt.Println("Metode pembayaran: Cash")
	case 2:
		fmt.Println("Metode pembayaran: Transfer")
	case 3:
		fmt.Println("Metode pembayaran: E-Wallet")
	default:
		fmt.Println("Metode tidak dikenal")
	}

	// 17. Fungsi multiple return → cicilan
	var cicilan int
	fmt.Print("\nMasukkan jumlah cicilan: ")
	fmt.Scanln(&cicilan)

	angsuran, err := bagi(totalAkhir, cicilan)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cicilan per bulan: Rp.%d\n", angsuran)
	}

	// Ringkasan belanja
	fmt.Println("\n=== Ringkasan Belanja ===")
	fmt.Println("Pembeli:", namaPembeli)
	fmt.Println("Kasir:", namaKasir)
	fmt.Println("Barang yang dibeli:")
	for _, b := range barang {
		fmt.Printf("- %s : Rp.%d\n", b, harga[b])
	}
	fmt.Println("Total Belanja:", totalBelanja)
	fmt.Println("Diskon:", potongan)
	fmt.Println("PPN:", ppn)
	fmt.Println("Total Bayar:", totalAkhir)
}
