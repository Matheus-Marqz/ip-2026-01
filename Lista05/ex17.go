package main

import "fmt"

func ehPrimo(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var v [10]int
	encontrou := false

	fmt.Println("Digite 10 numeros inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Println("Numeros primos e suas posicoes:")
	for i := 0; i < 10; i++ {
		if ehPrimo(v[i]) {
			fmt.Printf("Numero: %d | Posicao: %d\n", v[i], i)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Nao ha numeros primos no vetor.")
	}
}
