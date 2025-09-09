package main

import "fmt"

func main() {

	var inputUser string
	var inputPass string

	var username = "saya"
	var password = "123456"
	fmt.Println("Masukkan username: ")
	fmt.Scanln(&inputUser)
	fmt.Println("Masukkan password: ")
	fmt.Scanln(&inputPass)

	if inputUser == username && inputPass == password {
		fmt.Println("Login berhasil")
	} else {
		fmt.Println("Login Gagal")
	}
}
