package main

import "fmt"

func main() {
	var v1 [10]int
	var v2 [5]int
	var pares []int
	var impares []int
	somaV2 := 0

	fmt.Println("Digite os 10 elementos do primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v1[i])
	}

	fmt.Println("Digite os 5 elementos do segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&v2[i])
		somaV2 += v2[i]
	}

	for i := 0; i < 10; i++ {
		if v1[i]%2 == 0 {
			pares = append(pares, v1[i]+somaV2)
		} else {
			impares = append(impares, v1[i]+somaV2)
		}
	}

	fmt.Println("Primeiro vetor resultante (pares):", pares)
	fmt.Println("Segundo vetor resultante (impares):", impares)
}
