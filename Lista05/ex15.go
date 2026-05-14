package main

import "fmt"

func main() {
	var original [30]int
	var gerado [30]int

	fmt.Println("Digite 30 numeros inteiros:")
	for i := 0; i < 30; i++ {
		fmt.Scan(&original[i])
	}

	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			gerado[i] = original[i] * 2
		} else {
			gerado[i] = original[i] * 3
		}
	}

	fmt.Println("Vetor gerado:")
	for i := 0; i < 30; i++ {
		fmt.Printf("%d ", gerado[i])
	}
	fmt.Println()
}
