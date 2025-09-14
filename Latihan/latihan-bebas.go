package main

import "fmt"

func main() {

	var jumlah int
	fmt.Print("Enter jumlah: ")
	fmt.Scanln(&jumlah)

	barang := make([]string, 0)

	harga := make(map[string]int)

	for i := 0; i < jumlah; i++ {
		var name string
		var hrg int
		fmt.Print("Enter nama: ")
		fmt.Scanln(&name)
		fmt.Print("Enter harga: ")
		fmt.Scanln(&hrg)
		barang = append(barang, name)
		harga[name] = hrg
	}

	fmt.Println("\n=== Daftar Belanja ===")
	total := 0
	for _, name := range barang {
		fmt.Printf("%s: Rp.%d\n", name, harga[name])
		total += harga[name]
	}

	fmt.Printf("\n=== Total Belanja === Rp.%d\n", total)
}
