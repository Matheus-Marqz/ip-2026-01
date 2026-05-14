package main

import (
	"fmt"
	"math"
)

func main() {
	var entrada [15]int
	var resultado [15]float64

	fmt.Println("Digite 15 numeros inteiros:")
	for i := 0; i < 15; i++ {
		fmt.Scan(&entrada[i])

		if entrada[i] < 0 {
			resultado[i] = -1
		} else {
			resultado[i] = math.Sqrt(float64(entrada[i]))
		}
	}

	fmt.Println("Valores armazenados no vetor resultado:")
	for i := 0; i < 15; i++ {
		fmt.Printf("%.2f ", resultado[i])
	}
	fmt.Println()
}
