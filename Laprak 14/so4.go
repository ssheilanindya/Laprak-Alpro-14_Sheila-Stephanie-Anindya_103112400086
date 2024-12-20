package main

import "fmt"

func main() {
	var pita, kumpulanPita string
	var n, totalBunga int

	fmt.Scan(&n)
	if n < 0 {
		fmt.Print("Tidak valid")

	} else {
		for i := 1; i <= n; i++ {
			fmt.Print("Bunga : ")
			fmt.Scan(&pita)
			if pita == "DONE" {
				break
			}
			kumpulanPita = kumpulanPita + (pita + "-")
			totalBunga++
		}
	}
	fmt.Print("pita = ", kumpulanPita)
	fmt.Println("Banyak bunga = ", totalBunga)
}
