// membuat inputan data mahasiswa

package main

import "fmt"

// deklarasi variabel
var npm, nama, alamat, jurusan string

func main() {
	fmt.Print("Masukkan npm     : ")
	fmt.Scan(&npm)
	fmt.Print("Masukkan nama    : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan alamat  : ")
	fmt.Scan(&alamat)
	fmt.Print("Masukkan jurusan : ")
	fmt.Scan(&jurusan)

	fmt.Println("Npm     : ", npm)
	fmt.Println("Nama    : ", nama)
	fmt.Println("Alamat  : ", alamat)
	fmt.Println("Jurusan : ", jurusan)
}

// Pada program ini, terdapat beberapa variabel yang dideklarasikan di awal, yaitu npm, nama, alamat,
// dan jurusan. Kemudian, di dalam blok main(), program akan meminta input dari user untuk setiap
// variabel tersebut menggunakan perintah fmt.Scan(). Setelah itu, program akan menampilkan input
// yang telah dimasukkan oleh user. Dengan demikian, logika dari program ini adalah untuk membuat
// inputan data mahasiswa dengan menggunakan variabel npm, nama, alamat, dan jurusan, dan
// menampilkan hasil input tersebut ke layar.
