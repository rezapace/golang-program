package main

import "fmt"

func main() {
	// Deklarasi variabel nilai
	var nilai int
	fmt.Print("Masukkan nilai Anda: ")
	fmt.Scan(&nilai)

	// Cek apakah nilai lebih besar atau sama dengan 50
	if nilai >= 50 {
		fmt.Println("Nilai Anda:", nilai)
	} else {
		fmt.Println("Maaf, Anda tidak lulus")
	}
}

// program dengan output:
//     - nilai anda adalah 50
//     - maaf anda tidak lulus

// Program ini akan meminta user untuk memasukkan nilai, kemudian akan mengecek apakah nilai
// tersebut lebih besar atau sama dengan 50. Jika ya, maka akan menampilkan pesan "Nilai Anda:
// <nilai>", di mana <nilai> adalah nilai yang dimasukkan oleh user; sebaliknya, akan menampilkan pesan
// "Maaf, Anda tidak lulus".
