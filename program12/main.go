// membuat segitiga siku-siku dengan bintang
package main

import "fmt"

func main() {
	var tinggi int
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
