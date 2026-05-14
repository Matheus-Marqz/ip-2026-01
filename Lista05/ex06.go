package main

import "fmt"

func main() {
	var v [100]int
	valor := 100

	for i := 0; i < 100; i++ {
		v[i] = valor
		valor--
	}

	fmt.Println("Vetor em ordem decrescente de 100 a 1:")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", v[i])
	}
	fmt.Println()
}
