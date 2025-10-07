package main

import "fmt"

func SapaHello() {
	fmt.Println("------------------------")
	fmt.Println("Hallo World")
	fmt.Println("------------------------")
}

func Pembatas() {
	fmt.Println("------------------------")
	fmt.Println("Batas")
	fmt.Println("------------------------")
}

func main() {
	// ini sigle func
	fmt.Println("Ini Jalan di Go")

	// ini multifunction
	SapaHello()

	Pembatas()

	fmt.Println("Type Data")
	// Belajar Variabel dan Type Data
	// cara 1
	var Nama1 string = "wicaksono"
	fmt.Println("Nama Saya ", Nama1)

	// Cara 2
	Nama1 = "sono"
	fmt.Println("Nama Satu di ganti = ", Nama1)

	// ini coba ke dua

	var Nama2 string
	Nama2 = "Wicak"

	fmt.Println("Nama Saya menjadi = ", Nama2)

	// cara 3
	Umur := 30
	statusmenikah := false
	fmt.Println("Nama adalah", Nama2, "dan umur adalah", Umur, statusmenikah)

	// constanta (nilai TIdak dapat di rubah)
	const Phi = 3.14

	Pembatas()
	// Operator arimatika
	a := 5
	b := 10
	c := a * b
	fmt.Println("Perkalian = ", c)

	// Operator Perbandingan

}
