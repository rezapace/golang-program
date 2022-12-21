// menghitung luas persegi
package main

import "fmt"

func main() {
	var sisi int
	fmt.Print("Masukkan sisi persegi: ")
	fmt.Scan(&sisi)

	luas := sisi * sisi
	fmt.Println("Luas persegi adalah", luas)
}
