package main

import "fmt"

//var angka = []int{9, 9, 12, 7}
//
//func jumlah(angka []int) int {
//	total := 0
//	for _, v := range angka {
//		total += v
//	}
//	return total
//}
//
//func main() {
//	totalHasil := jumlah(angka)
//	fmt.Println(totalHasil)
//}

func main() {
	var a int
	fmt.Println("Jumlah buku: ")
	fmt.Scan(&a)
	// buat slice kosong
	var buku []string
	for i := 0; i < a; i++ {
		var b string
		fmt.Printf("Buku ke-%d: ", i+1)
		fmt.Scan(&b)
		buku = append(buku, b)
	}
	fmt.Println("Daftar buku:", buku)
}
func main() {
	var a int
	fmt.Print("ada berapa angka?: ")
	fmt.Scan(&a)

	// buat slice dengan panjang a
	angka := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Printf("Masukkan angka ke-%d:", i+1)
		fmt.Scan(&angka[i]) // langsung masuk ke posisi slice
	}
	fmt.Println("Daftar angka:", angka)
}
