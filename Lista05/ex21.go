package main

import "fmt"

func main() {
	var codigo int
	var v [10]float64

	fmt.Println("Digite 10 numeros reais:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Println("Digite o codigo (0 para terminar, 1 para ordem direta, 2 para ordem inversa):")
	fmt.Scan(&codigo)

	if codigo == 0 {
		fmt.Println("Programa finalizado.")
		return
	}

	if codigo == 1 {
		fmt.Println("Vetor na ordem direta:")
		for i := 0; i < 10; i++ {
			fmt.Printf("%.2f ", v[i])
		}
		fmt.Println()
	} else if codigo == 2 {
		fmt.Println("Vetor na ordem inversa:")
		for i := 9; i >= 0; i-- {
			fmt.Printf("%.2f ", v[i])
		}
		fmt.Println()
	} else {
		fmt.Println("Codigo invalido.")
	}
}
