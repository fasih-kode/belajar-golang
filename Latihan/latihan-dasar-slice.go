package main

import "fmt"

var angka = []int{}

func jumlah(angka []int) int {
	total := 0
	for _, v := range angka {
		total += v
	}
	return total
}

func main() {
	totalHasil := jumlah(angka)
	fmt.Println(totalHasil)
}
