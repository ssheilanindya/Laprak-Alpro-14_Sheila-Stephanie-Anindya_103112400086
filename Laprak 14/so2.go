package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a == 1 {
		fmt.Println("bukan prima")
		return
	}

	prima := true
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			prima = false
			break
		}
	}

	if prima {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}
