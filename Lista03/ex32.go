package main

import "fmt"

func main() {
	var n1, n2, resultado int
	var negativo bool

	fmt.Print("Digite o primeiro numero: ")
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo numero: ")
	fmt.Scan(&n2)

	if n1 < 0 {
		n1 = -n1
		negativo = !negativo
	}

	if n2 < 0 {
		n2 = -n2
		negativo = !negativo
	}

	for i := 1; i <= n2; i++ {
		resultado = resultado + n1
	}

	if negativo {
		resultado = -resultado
	}

	fmt.Println("Resultado:", resultado)
}
