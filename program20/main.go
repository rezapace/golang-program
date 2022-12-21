// menghitung luas permukaan kubus

package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Print("Masukkan sisi: ")
	fmt.Scan(&s)

	luas := 6 * math.Pow(s, 2)
	fmt.Println("Luas permukaan kubus adalah", luas)
}
