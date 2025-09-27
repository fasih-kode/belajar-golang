package main

import "fmt"

func main() {
	angka := make([]int, 3)

	for i := 0; i < 3; i++ {
		fmt.Printf("Angka ke-%v: ", i+1)
		fmt.Scan(&angka[i])
	}

	// cetak isi slice
	for i, v := range angka {
		fmt.Printf("Indeks ke-[%v] = %v\n", i, v)
	}

	// jumlah isi slice
	total := 0
	for _, v := range angka {
		total += v
	}
	fmt.Printf("Jumlah semuanya %v\n: ", total)
}
