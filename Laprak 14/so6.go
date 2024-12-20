package main

import "fmt"

func main() {
	var k int
	var penyebut, pembilang, hasil float64

	fmt.Print(" k : ")
	fmt.Scan(&k)

	for i := 0; i <= k; i++ {
		pembilang = (4*float64(k) + 2) * (4*float64(k) + 2)
		penyebut = (4*float64(k) + 1) * (4*float64(k) + 3)

		hasil = pembilang / penyebut
	}
	fmt.Printf("f(%d) = %.10f\n", k, hasil)
}
