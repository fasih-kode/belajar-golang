package greet

import "fmt"

func Hello() {
	var nama string
	fmt.Println("Hello dari fungsi yang diekspor")
	fmt.Println("Siapa nama anda?")
	fmt.Scan(&nama)
	fmt.Println("Nama anda bagus sekali")
}

//func hello() {
//	fmt.Println("Hello dari fungsi yang tidak di ekspor")
//}

// func harus kapital jika mau buat package
