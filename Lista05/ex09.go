package main

import "fmt"

func main() {
	var alturas [10]float64
	soma := 0.0
	media := 0.0

	fmt.Println("Digite a altura de 10 atletas:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&alturas[i])
		soma += alturas[i]
	}

	media = soma / 10.0
	fmt.Printf("Media das alturas: %.2f\n", media)

	fmt.Println("Alturas acima da media:")
	for i := 0; i < 10; i++ {
		if alturas[i] > media {
			fmt.Printf("%.2f\n", alturas[i])
		}
	}
}
