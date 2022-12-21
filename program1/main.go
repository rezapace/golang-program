package main

import "fmt"

func main() {
	// Deklarasi variabel mutu
	var mutu int
	fmt.Print("Masukkan nilai mutu: ")
	fmt.Scan(&mutu)

	// Percabangan if-else
	if mutu <= 40 {
		fmt.Println("Nilai mutu Anda: D")
	} else if mutu <= 60 {
		fmt.Println("Nilai mutu Anda: C")
	} else if mutu <= 80 {
		fmt.Println("Nilai mutu Anda: B")
	} else {
		fmt.Println("Nilai mutu Anda: A")
	}
}

// percabangan Golang dengan kondisi if-else
//     - Jika mutu kurang dari sama dengan 40 maka nilai D,
//     - jika mutu kurang dari 61 maka nilai C,
//     - jika nilai kurang dari 81 maka nilai B,
//     - selain itu nilai A.

// Program ini akan meminta user untuk memasukkan nilai mutu, kemudian akan mengecek nilai mutu
// tersebut berdasarkan percabangan yang telah ditentukan. Jika mutu kurang dari atau sama dengan 40,
// maka akan menampilkan nilai D; jika mutu kurang dari 61, maka akan menampilkan nilai C; jika mutu
// kurang dari 81, maka akan menampilkan nilai B; selain itu, akan menampilkan nilai A.
