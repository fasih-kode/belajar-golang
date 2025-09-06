package main

import "fmt"

func main() {

	var nilai int
	fmt.Println("Masukkan nilai kamu: ")
	fmt.Scanln(&nilai)

	if nilai < 0 || nilai > 100 {
		fmt.Println("Nilai tidak valid")
	} else if nilai >= 90 {
		fmt.Println("Nilai kamu A")
	} else if nilai >= 80 {
		fmt.Println("Nilai kamu B")
	} else if nilai >= 70 {
		fmt.Println("Nilai kamu C")
	} else {
		fmt.Println("Nilai kamu D")
	}
}
