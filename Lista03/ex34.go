package main

import "fmt"

func main() {
	var n1, n2, mmc, maior int

	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)
	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	if n1 > n2 {
		maior = n1
	} else {
		maior = n2
	}

	mmc = maior

	for mmc%n1 != 0 || mmc%n2 != 0 {
		mmc++
	}

	fmt.Println("MMC:", mmc)
}
