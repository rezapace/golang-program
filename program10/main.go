package main

import "fmt"

func main() {
	// Membuat array dengan nama-nama hari
	hari := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	// Perulangan untuk menampilkan nama hari
	for i := 0; i < len(hari); i++ {
		fmt.Println("Hari ke", i+1, "adalah hari", hari[i])
	}
}

// program 10
// Buat sebuah array berisi nama-nama hari
// dari senin hingga minggu
// dengan contoh
//     - hari ke 1 adalah hari senin
//     - hari ke 2 adalah hari selasa

// Program ini akan membuat sebuah array dengan nama hari yang berisi nama-nama hari dari senin
// hingga minggu. Kemudian, program akan melakukan perulangan untuk menampilkan nama hari pada
// setiap index array dengan menggunakan fungsi fmt.Println(). Setelah seluruh hari telah ditampilkan,
// program akan selesai.
