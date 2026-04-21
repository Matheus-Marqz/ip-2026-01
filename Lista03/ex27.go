package main

import (
	"fmt"
	"math"
)

func main() {
	var x, cosSerie, cosFuncao, diferenca float64
	var sinal float64 = 1

	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)

	for i := 0; i < 20; i++ {
		expoente := 2 * i
		fatorial := 1.0
		potencia := 1.0

		for j := 1; j <= expoente; j++ {
			fatorial = fatorial * float64(j)
			potencia = potencia * x
		}

		cosSerie = cosSerie + sinal*(potencia/fatorial)
		sinal = sinal * -1
	}

	cosFuncao = math.Cos(x)
	diferenca = cosSerie - cosFuncao

	fmt.Println("Cosseno pela serie:", cosSerie)
	fmt.Println("Diferenca:", diferenca)
}
