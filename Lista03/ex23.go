package main

import "fmt"

func main() {
	var n int
	var s float64
	var numerador float64 = 1000
	var sinal float64 = 1

	fmt.Print("Digite a quantidade de termos: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Valor invalido")
		return
	}

	for i := 1; i <= n; i++ {
		s = s + sinal*(numerador/float64(i))
		numerador = numerador - 3
		sinal = sinal * -1
	}

	fmt.Println("S =", s)
}
