package main

import "fmt"

// Fungsi untuk menghitung luas persegi
func hitungLuasPersegi(panjang, lebar float64) float64 {
	return panjang * lebar
}

func main() {
	// Deklarasi variabel panjang dan lebar
	var panjang, lebar float64
	fmt.Print("Masukkan panjang persegi: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar persegi: ")
	fmt.Scan(&lebar)

	// Hitung luas persegi
	luas := hitungLuasPersegi(panjang, lebar)

	// Tampilkan hasil
	fmt.Println("Luas persegi:", luas)
}

// program 7
// program menghitung Luas Persegi dengan parameter

//     - Panjang dan Lebar

// Fungsi hitungLuasPersegi() akan menerima dua parameter, yaitu panjang dan lebar, lalu menghitung
// dan mengembalikan luas persegi berdasarkan rumus panjang * lebar. Fungsi ini kemudian akan
// dipanggil dari main() untuk menghitung luas persegi dan menampilkan hasilnya ke layar.
