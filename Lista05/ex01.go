package main

import "fmt"

func main() {
	var v [10]int
	encontrou := false

	fmt.Println("Digite 10 numeros inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Println("Numeros superiores a 50 e suas posicoes:")
	for i := 0; i < 10; i++ {
		if v[i] > 50 {
			fmt.Printf("Numero: %d | Posicao: %d\n", v[i], i)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Nao existe nenhum numero superior a 50.")
	}
}
