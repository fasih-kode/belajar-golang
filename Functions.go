package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int, z int) int {
	return x * y * z
}

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

func main() {
	fmt.Println(add(20, 90))
	fmt.Println(multiply(76, 21, 0))
	fmt.Println(swap("Boss", "Siap"))
	sayHeloo()
	a, b := greet("Fasih", 20)
	fmt.Println("Halo nama saya", a, "Umur saya", b, "Tahun")

	var x, y int
	fmt.Println("Masukkan angka pertama: ")
	fmt.Scan(&x)

	fmt.Println("Masukkan angka kedua: ")
	fmt.Scan(&y)

	hasil := add(x, y)
	fmt.printf("Hasil penjumlahan %d + %d = %d\n", x, y, hasil)
}
