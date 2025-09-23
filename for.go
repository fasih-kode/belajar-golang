package main

import "fmt"

func main() {
	// menjumlahkan hasil dari perulangan
	var angka = 2
	for i := 0; i < 10; i++ {
		angka += i
	}

	// for biasa → sering dipakai untuk angka (counter loop)
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// FOR RANGE

	// for range → dipakai untuk koleksi (termasuk string)
	word := "halo"
	for i, v := range word {
		fmt.Println("Index", i, "Char", string(v))
	}

	// Looping slice (list angka/teks)
	nums := []int{10, 20, 30, 40, 50}

	for idx, val := range nums {
		fmt.Println("Index", idx, "Val", val)
	}

	// Looping map (key-value)
	m := map[string]int{"a": 1, "b": 2}

	for key, val := range m {
		fmt.Println("Key:", key, "Value:", val)
	}

	// Ini adalah loop for, tapi berbeda dari sebelumnya:
	//
	//Bagian init dan post dikosongkan (; ;).
	//
	//Yang tersisa hanya kondisi: sum < 1000.
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	// Artinya:
	//
	//Selama sum masih kurang dari 1000, perulangan terus berjalan.
	//
	//Di dalam loop: sum += sum artinya sum = sum + sum (menggandakan nilainya).
	//
	//📌 Jadi alurnya seperti ini:
	//
	//Awal: sum = 1
	//
	//Putaran 1 → sum = 2
	//
	//Putaran 2 → sum = 4
	//
	//Putaran 3 → sum = 8
	//
	//Putaran 4 → sum = 16
	//
	//Putaran 5 → sum = 32
	//
	//Putaran 6 → sum = 64
	//
	//Putaran 7 → sum = 128
	//
	//Putaran 8 → sum = 256
	//
	//Putaran 9 → sum = 512
	//
	//Putaran 10 → sum = 1024 → kondisi sum < 1000 sudah salah, loop berhenti.

	fmt.Println(angka)
	fmt.Println(sum)
	//for {
	//	fmt.Print("-->")
	//}
}
