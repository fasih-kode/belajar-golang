package main

import "fmt"

//func add(x int, y int) int {
//	return x + y
//}

//func multiply(x int, y int, z int) int {
//	return x * y * z
//}

func swap(x, y string) (string, string) {
	return y, x
}

func sayHeloo() {
	fmt.Println("Hello Go!")
}

func greet(name string, age int) (string, int) {
	return name, age
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (float64, int) {
	return float64(a) / float64(b), a % b
}

func splitEvenOdd(sum int) (even, odd int) {
	even = sum / 2
	odd = sum - even
	return
}

// konversi waktu
func jamWaktu(jumlahDetik int) (jam, menit, detik int, hari int) {
	jam = jumlahDetik / 3600
	menit = (jumlahDetik % 3600) / 60
	detik = jumlahDetik % 60
	hari = jam / 24
	return
}

func gantiHari(jumlahHari int) (jam, menit, detik int) {
	jam = jumlahHari * 24
	menit = jam * 60
	detik = menit * 60
	return
}

func main() {
	fmt.Println(add(20, 90))
	//fmt.Println(multiply(76, 21, 0))
	fmt.Println(swap("Boss", "Siap"))
	sayHeloo()
	a, b := greet("Fasih", 20)
	fmt.Println("Halo nama saya", a, "Umur saya", b, "Tahun")

	var x, y int
	fmt.Println("Masukkan angka pertama: ")
	fmt.Scan(&x)

	fmt.Println("Masukkan angka kedua: ")
	fmt.Scan(&y)

	var det int
	fmt.Println("Masukkan jumlah detik: ")
	fmt.Scan(&det)

	var har int
	fmt.Println("Masukkan jumlah hari: ")
	fmt.Scan(&har)

	hasilTambah := add(x, y)
	hasilKurang := subtract(x, y)
	hasilKali := multiply(x, y)
	hasilBagi, sisa := divide(x, y)
	hasilEven, hasilOdd := splitEvenOdd(y)
	jam, menit, detik, hari := jamWaktu(det)
	j, m, d := gantiHari(har)

	fmt.Printf("Hasil penjumlahan %d + %d = %d\n", x, y, hasilTambah)
	fmt.Printf("Hasil pengurangan %d - %d = %d\n", x, y, hasilKurang)
	fmt.Printf("Hasil perkalian %d * %d = %d\n", x, y, hasilKali)
	fmt.Printf("Hasil pembagian %d / %d = %.2f sisa %d\n", x, y, hasilBagi, sisa)
	fmt.Printf("Hasil split dari %d adalah genap: %d dan ganjil: %d\n", y, hasilEven, hasilOdd)
	fmt.Printf("Hasil konversi %d detik adalah %d hari, %d jam, %d menit, %d detik\n", det, hari, jam, menit, detik)
	fmt.Printf("Hasil konversi hari %d adalah %d jam, %d menit, %d detik\n", har, j, m, d)

}
