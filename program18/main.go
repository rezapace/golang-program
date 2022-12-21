// menghitung masa jenis
package main

import "fmt"

func main() {
	var massa, volume int
	fmt.Print("Masukkan massa: ")
	fmt.Scan(&massa)
	fmt.Print("Masukkan volume: ")
	fmt.Scan(&volume)

	masaJenis := massa / volume
	fmt.Println("Masa jenis adalah", masaJenis)
}
