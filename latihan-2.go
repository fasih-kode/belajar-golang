package main

import "fmt"

func hitungLuasPersegi(s1, s2 float64) float64 {
	return s1 * s2
}

func hitungKelilingPersegi(s1, s2, s3, s4 float64) float64 {
	return s1 + s2 + s3 + s4
}

func main() {
	var sisi1, sisi2, sisi3, sisi4 float64

	// hitung luas persegi
	fmt.Println("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi1)
	fmt.Println("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi2)

	luas := hitungLuasPersegi(sisi1, sisi2)
	fmt.Printf("Luas persegi dengan sisi %.2f dan sisi %.2f adalah %.2f\n", sisi1, sisi2, luas)

	// hitung keliling persegi
	fmt.Println("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi1)
	fmt.Println("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi2)
	fmt.Println("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi3)
	fmt.Println("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi4)

	keliling := hitungKelilingPersegi(sisi1, sisi2, sisi3, sisi4)
	fmt.Printf("Keliling persegi dengan sisi1 %.2f sisi2 %.2f sisi3 %.2f sisi4 %.2f adalah %.2f\n", sisi1, sisi2, sisi3, sisi4, keliling)
}
