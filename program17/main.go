// menghitung  sudut sin, cos, tan
package main

import (
	"fmt"
	"math"
)

func main() {
	var sudut float64
	fmt.Print("Masukkan sudut: ")
	fmt.Scan(&sudut)

	sin := math.Sin(sudut)
	cos := math.Cos(sudut)
	tan := math.Tan(sudut)

	fmt.Println("Nilai sin adalah", sin)
	fmt.Println("Nilai cos adalah", cos)
	fmt.Println("Nilai tan adalah", tan)
}
