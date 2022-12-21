// Buatlah program golang dengan output berupa angka genap 1-75 dengan
package main

import "fmt"

func main() {
	// Perulangan untuk mengecek setiap bilangan dari 1-75
	for i := 1; i <= 75; i++ {
		// Cek apakah bilangan tersebut genap
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// program 9
// program dengan output berupa angka
//     - genap 1-75

// Berikut adalah logika program yang dapat menampilkan output berupa angka genap dari bilangan 1-75:

// 1.Pertama, program akan mengecek setiap bilangan dari 1-75 dengan menggunakan perulangan for.
// 2.Selanjutnya, dalam setiap iterasi perulangan, program akan mengecek apakah bilangan tersebut
// genap atau tidak dengan menggunakan operator modulus % (menghasilkan sisa bagi dari pembagian).
// Jika bilangan tersebut genap (hasil modulus 0), maka program akan menampilkan bilangan tersebut
// ke layar dengan menggunakan fungsi fmt.Println().
// 3.Setelah seluruh bilangan dari 1-75 telah dicek, program akan selesai dan menampilkan hasilnya.
