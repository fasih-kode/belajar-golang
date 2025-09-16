package main

import "fmt"

// fungsi tanpa parameter & tanpa return
func sapa() {
	fmt.Println("Halo, selamat datang!")
}

// fungsi dengan parameter
func tambah(a int, b int) {
	fmt.Println("Hasil:", a+b)
}

// fungsi dengan return
func kali(a int, b int) int {
	return a * b
}

// fungsi dengan multiple return
func bagi(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa dibagi nol")
	}
	return a / b, nil
}

// fungsi utama wajib ada
func main() {
	sapa()

	tambah(2, 3)

	hasilKali := kali(4, 5)
	fmt.Println("Hasil Kali", hasilKali)

	hasilBagi, err := bagi(10, 0)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Hasil Bagi", hasilBagi)
	}
}
