package main

import "fmt"

func main() {
	var b [100]float64
	soma := 0.0

	fmt.Println("Digite 100 valores numericos:")
	for i := 0; i < 100; i++ {
		fmt.Scan(&b[i])
	}

	for i := 0; i < 50; i++ {
		diferenca := b[i] - b[99-i]
		soma += diferenca * diferenca * diferenca
	}

	fmt.Printf("Valor de S = %.2f\n", soma)
}
