package main

import "fmt"

func main() {
	var berat1, berat2 float64

	for {
		fmt.Print("Berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai")
			break
		}
		if berat1+berat2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		if berat1-berat2 >= 9 || berat2-berat1 >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
