package main

import "fmt"

func main() {
	var v [10]int
	menorPos := 0

	fmt.Println("Digite 10 numeros inteiros diferentes:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v[i])
	}

	for i := 1; i < 10; i++ {
		if v[i] < v[menorPos] {
			menorPos = i
		}
	}

	fmt.Printf("O menor elemento do vetor e %d e sua posicao dentro do vetor e: %d\n", v[menorPos], menorPos)
}
