package main

import "fmt"

func main() {
	var graos uint64 = 1
	var total uint64

	for i := 1; i <= 64; i++ {
		total = total + graos
		graos = graos * 2
	}

	fmt.Println("Total de graos:", total)
}
