// membuat segitiga full dengan bintang

package main

import "fmt"

func main() {
	var tinggi int
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	for brs := 1; brs <= tinggi; brs++ {
		//melakukan pengulangan bintang(*) sampai 30

		for spasi := 30; spasi >= brs; spasi-- {
			fmt.Printf(" ")
		}

		for klm := 1; klm <= brs; klm++ {
			fmt.Printf("*")

		}
		for klm := (brs - 1); klm >= 1; klm-- {
			fmt.Printf("*")

		}
		fmt.Printf("\n")
	}
}
