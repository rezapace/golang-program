package main

import "fmt"

// Fungsi untuk menghitung luas segitiga
func hitungLuasSegitiga(alas, tinggi float64) float64 {
	return alas * tinggi / 2
}

func main() {
	// Deklarasi variabel alas dan tinggi
	var alas, tinggi float64
	fmt.Print("Masukkan alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	// Hitung luas segitiga
	luas := hitungLuasSegitiga(alas, tinggi)

	// Tampilkan hasil
	fmt.Println("Luas segitiga:", luas)
}

// Kalkulator/Fungsi untuk menghitung luas Segitiga
// dengan parameter

// Fungsi hitungLuasSegitiga() akan menerima dua parameter, yaitu alas dan tinggi, lalu menghitung dan
// mengembalikan luas segitiga berdasarkan rumus alas * tinggi / 2. Fungsi ini kemudian akan dipanggil
// dari main() untuk menghitung luas segitiga dan menampilkan hasilnya ke layar.
