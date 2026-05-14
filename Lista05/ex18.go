package main

import "fmt"

func main() {
	var v [10]int

	fmt.Println("Digite 10 numeros inteiros em ordem crescente:")
	for i := 0; i < 10; i++ {
		for {
			fmt.Scan(&v[i])

			if i == 0 || v[i] >= v[i-1] {
				break
			}

			fmt.Println("Valor fora da ordem crescente. Digite novamente:")
		}
	}

	fmt.Println("Vetor lido em ordem crescente:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", v[i])
	}
	fmt.Println()
}
