package main

import "fmt"

func main() {
	var b, n, resultado int

	fmt.Print("Digite o valor de b: ")
	fmt.Scan(&b)
	fmt.Print("Digite o valor de n: ")
	fmt.Scan(&n)

	resultado = 1

	for i := 1; i <= n; i++ {
		resultado = resultado * b
	}

	fmt.Println("Resultado:", resultado)
}
