package main

import "fmt"

func main() {
	var preco, ingressos, lucro, maiorLucro, melhorPreco float64
	var melhorQtd int
	var qtd int = 130

	fmt.Println("Preco\tIngressos\tLucro")

	for preco = 6.0; preco >= 1.0; preco = preco - 0.6 {
		ingressos = float64(qtd)
		lucro = preco*ingressos - 300

		fmt.Printf("%.2f\t%d\t\t%.2f\n", preco, qtd, lucro)

		if qtd == 130 || lucro > maiorLucro {
			maiorLucro = lucro
			melhorPreco = preco
			melhorQtd = qtd
		}

		qtd = qtd + 30
	}

	fmt.Println("Maior lucro:", maiorLucro)
	fmt.Println("Melhor preco:", melhorPreco)
	fmt.Println("Quantidade de ingressos:", melhorQtd)
}
