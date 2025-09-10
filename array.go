package main

import "fmt"

func main() {

	var nilai [5]int
	nilai[0] = 60
	nilai[1] = 80
	nilai[2] = 90
	nilai[3] = 75
	nilai[4] = 50
	fmt.Println(nilai)

	var hari [7]string
	hari[0] = "Senin"
	hari[1] = "Selasa"
	hari[2] = "Rabu"
	hari[3] = "Kamis"
	hari[4] = "Jumat"
	hari[5] = "Sabtu"
	hari[6] = "Ahad"
	fmt.Println("Hari ke-2:", hari[1])
	fmt.Println("Hari ke-5:", hari[4])

	var angka [4]int
	angka[0] = 32
	angka[1] = 20
	angka[2] = 10
	angka[3] = 50
	fmt.Println("jumlah semuanya:", angka[0]+angka[1]+angka[2]+angka[3])
}
