// mengitung 1- 1juta dan hitung berapa lama waktu yang dibutuhkan dalam satuan second

package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	var start = time.Now()
	for i = 1; i <= 1000000; i++ {
	}
	var end = time.Now()
	var elapsed = end.Sub(start)
	fmt.Println("Waktu yang dibutuhkan adalah", elapsed)
}
