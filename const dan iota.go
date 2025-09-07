package main

import "fmt"

const (
	Senin = iota + 1
	Selasa
	Rabu
	Kamis
	Jumat
	Sabtu
	Ahad
)

func main() {
	fmt.Println(Senin, Selasa, Rabu, Kamis, Jumat, Sabtu, Ahad)
}
