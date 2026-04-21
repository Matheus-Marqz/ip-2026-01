package main

import "fmt"

func main() {
	var n, soma int

	fmt.Print("Digite um numero inteiro positivo: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		soma = soma + i
	}

	fmt.Println("Soma:", soma)
}
