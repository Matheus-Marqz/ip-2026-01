package main

import "fmt"

func main() {
	var n int
	var fatorial uint64 = 1

	fmt.Print("Digite um numero inteiro: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Valor invalido")
		return
	}

	for i := 2; i <= n; i++ {
		fatorial = fatorial * uint64(i)
	}

	fmt.Println("Fatorial:", fatorial)
}
