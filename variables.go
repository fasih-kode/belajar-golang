package main

import "fmt"

// --- Var di luar fungsi (package level) ---
var pesanGlobal string = "Saya Global"
var counter int

func main() {
	// --- Var di dalam fungsi ---
	var nama string
	nama = "fasih" // assigment (=) setelah var

	// --- := di dalam fungsi ---
	age := 20 // langsung deklarasi + isi
	pesan := "Halo!"

	// --- Assigment (=) mengganti nilai ---
	age = 21
	pesan = "Selamat datang!"
	nama = "Reang"

	// output
	fmt.Println(pesanGlobal, counter)
	fmt.Println(nama, age, pesan)
}
