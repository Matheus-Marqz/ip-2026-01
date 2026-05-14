package main

import "fmt"

func main() {
	var lancamentos [20]int
	var freq [7]int

	fmt.Println("Digite os 20 numeros sorteados no dado (1 a 6):")
	for i := 0; i < 20; i++ {
		for {
			fmt.Scan(&lancamentos[i])
			if lancamentos[i] >= 1 && lancamentos[i] <= 6 {
				break
			}
			fmt.Println("Valor invalido. Digite um numero de 1 a 6:")
		}
		freq[lancamentos[i]]++
	}

	fmt.Println("Numeros sorteados:")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", lancamentos[i])
	}
	fmt.Println()

	fmt.Println("Frequencia de cada face:")
	for i := 1; i <= 6; i++ {
		fmt.Printf("Face %d: %d vez(es)\n", i, freq[i])
	}
}
