package main

import (
	"fmt"
	"strings"
)

func main() {
	// Deklarasi variabel string
	var input string
	input = "Programmer Manlan"

	// Menghitung jumlah karakter dengan fungsi len()
	jumlahKarakter := len(input)

	// Menghitung jumlah karakter dengan fungsi Count()
	jumlahKarakter2 := strings.Count(input, "") - 1

	// Tampilkan hasil
	fmt.Println("Jumlah karakter menggunakan fungsi len():", jumlahKarakter)
	fmt.Println("Jumlah karakter menggunakan fungsi Count():", jumlahKarakter2)
}

// program menampilkan jumlah
//     - karakter dalam suatu string
//     - dengan inputan string “Programmer”

// Program ini akan meminta user untuk memasukkan sebuah string, kemudian akan menghitung jumlah
// karakter dalam string tersebut menggunakan fungsi len() dan strings.Count(). Fungsi len() akan
// mengembalikan jumlah elemen (karakter) dalam sebuah string, sedangkan fungsi strings.Count() akan
// mengembalikan jumlah karakter dalam string dengan menghitung jumlah substring yang ditentukan
// (dalam contoh ini, kita menggunakan string kosong sebagai substring yang akan dihitung). Hasil dari
// kedua fungsi tersebut kemudian akan ditampilkan ke layar.
