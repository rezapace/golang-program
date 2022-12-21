// menghitung luas persegi panjang
package main

import "fmt"

func main() {
	var panjang, lebar int
	fmt.Print("Masukkan panjang persegi panjang: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar persegi panjang: ")
	fmt.Scan(&lebar)

	luas := panjang * lebar
	fmt.Println("Luas persegi panjang adalah", luas)
}
