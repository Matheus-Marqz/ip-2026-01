package main

import "fmt"

func main() {
	var notas [15]int
	var freq [11]int

	fmt.Println("Digite as 15 notas (de 0 a 10):")
	for i := 0; i < 15; i++ {
		for {
			fmt.Scan(&notas[i])
			if notas[i] >= 0 && notas[i] <= 10 {
				break
			}
			fmt.Println("Nota invalida. Digite novamente:")
		}
		freq[notas[i]]++
	}

	fmt.Println("Nota | Frequencia Absoluta | Frequencia Relativa")
	for nota := 0; nota <= 10; nota++ {
		relativa := float64(freq[nota]) / 15.0
		fmt.Printf("%4d | %18d | %18.2f\n", nota, freq[nota], relativa)
	}
}
