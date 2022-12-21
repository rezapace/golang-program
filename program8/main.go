package main

import "fmt"

// Fungsi untuk menghitung luas lingkaran
func hitungLuasLingkaran(r float64) float64 {
	pi := 3.14
	return pi * r * r
}

func main() {
	// Deklarasi variabel jari-jari
	var r float64 = 7

	// Hitung luas lingkaran
	luas := hitungLuasLingkaran(r)

	// Tampilkan hasil
	fmt.Println("Luas lingkaran:", luas)
}

// program 8
// program untuk menghitung luas lingkaran dengan pi = 3,14 r = 7

// Fungsi hitungLuasLingkaran() akan menerima satu parameter, yaitu jari-jari r, lalu menghitung dan
// mengembalikan luas lingkaran berdasarkan rumus pi * r * r, di mana pi diisi dengan nilai 3.14. Fungsi
// ini kemudian akan dipanggil dari main() untuk menghitung luas lingkaran dan menampilkan hasilnya
// ke layar.
