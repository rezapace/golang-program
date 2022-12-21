package main

import "fmt"

func main() {
	// Deklarasi suku pertama dan suku kedua
	suku1 := 6
	suku2 := 9

	// Tampilkan suku pertama dan kedua
	fmt.Println(suku1)
	fmt.Println(suku2)

	// Perulangan untuk menghitung bilangan fibonacci selanjutnya
	for i := 1; i <= 3; i++ {
		// Hitung suku selanjutnya
		sukuSelanjutnya := suku1 + suku2

		// Tampilkan suku selanjutnya
		fmt.Println(sukuSelanjutnya)

		// Set suku1 menjadi suku2 dan suku2 menjadi suku selanjutnya
		suku1 = suku2
		suku2 = sukuSelanjutnya
	}
}

// program 11
// Buatlah program dengan output berupa 5 fibonacci number dengan

//     - suku pertama = 6
//     - suku ke kedua = 9

// Program ini akan membuat variabel suku1 dan suku2 untuk menyimpan suku pertama dan suku
// kedua dari deret fibonacci. Kemudian, program akan menampilkan nilai suku pertama dan kedua ke
// layar dengan menggunakan fungsi fmt.Println(). Selanjutnya, program akan melakukan perulangan
// untuk menghitung bilangan fibonacci selanjutnya sebanyak 3 kali. Pada setiap iterasi perulangan,
// program akan menghitung suku selanjutnya dengan menjumlahkan suku1 dan suku2, kemudian
// menampilkan suku selanjutnya ke layar. Setelah
