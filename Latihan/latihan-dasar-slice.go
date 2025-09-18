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
	fmt.Println("Jumlah siswa: ")
	fmt.Scanln(&a)

	// buat slice kosong
	var nama []string
	for i := 0; i < a; i++ {
		var b string
		fmt.Printf("Siswa ke%d: ", i+1)
		fmt.Scanln(&b)
		nama = append(nama, b)
	}
	fmt.Println("Daftar siswa:", nama)
}
