// kalkulator sederhana bahasa golang
package main

import "fmt"

func main() {
	// pendeklarasian variabel
	var angka1, angka2, tambah, kurang, kali, bagi int

	// input angka
	fmt.Print("Masukkan angka pertama : ")
	fmt.Scan(&angka1)
	fmt.Print("Masukkan angka kedua   : ")
	fmt.Scan(&angka2)

	// proses perhitungan
	tambah = angka1 + angka2
	kurang = angka1 - angka2
	kali = angka1 * angka2
	bagi = angka1 / angka2

	// output hasil
	fmt.Println(angka1, "+", angka2, " = ", tambah)
	fmt.Println(angka1, "-", angka2, " = ", kurang)
	fmt.Println(angka1, "*", angka2, " = ", kali)
	fmt.Println(angka1, "/", angka2, " = ", bagi)
}
