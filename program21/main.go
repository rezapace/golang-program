// program golang dengan output berupa angka ganjil 1-75

package main

import "fmt"

func main() {
	// Perulangan untuk mengecek setiap bilangan dari 1-75
	for i := 1; i <= 75; i++ {
		// Cek apakah bilangan tersebut ganjil
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

// Dari code di atas, program ini akan melakukan perulangan dari 1 hingga 75. Setiap bilangan yang
// didapat akan dicek apakah bilangan tersebut ganjil atau tidak dengan menggunakan operator modulo
// (%). Jika bilangan tersebut ganjil, maka akan dicetak ke layar menggunakan fmt.Println. Jadi, program
// ini akan mencetak semua bilangan ganjil dari 1 hingga 75 ke layar.
