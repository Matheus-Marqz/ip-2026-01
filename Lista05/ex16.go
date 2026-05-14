package main

import "fmt"

func main() {
	var idades [50]int
	moda := 0
	maiorFreq := 0

	fmt.Println("Digite 50 idades:")
	for i := 0; i < 50; i++ {
		fmt.Scan(&idades[i])
	}

	for i := 0; i < 50; i++ {
		cont := 0
		for j := 0; j < 50; j++ {
			if idades[j] == idades[i] {
				cont++
			}
		}

		if cont > maiorFreq {
			maiorFreq = cont
			moda = idades[i]
		}
	}

	fmt.Printf("Moda das idades: %d\n", moda)
	fmt.Printf("Quantidade de repeticoes da moda: %d\n", maiorFreq)
}
