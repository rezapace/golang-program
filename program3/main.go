package main

import "fmt"

func main() {
	// Perulangan untuk mengecek setiap bilangan dari 1-100
	for i := 1; i <= 100; i++ {
		// Cek apakah bilangan tersebut ganjil
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

// program Golang yang dapat menampilkan output berupa
//     - angka ganjil dari bilangan 1-100.

// Program ini akan melakukan perulangan dari bilangan 1-100, kemudian mengecek setiap bilangan
// tersebut apakah ganjil atau tidak. Jika bilangan tersebut ganjil, maka akan ditampilkan ke layar.
