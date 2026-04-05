package main

import "fmt"

func main() {
	var preco float64
	var codigo int
	var valorFinal float64

	fmt.Println("Digite o preco normal do produto:")
	fmt.Scan(&preco)

	fmt.Println("Escolha a condicao de pagamento:")
	fmt.Println("1 - A vista, dinheiro ou cheque (10% de desconto)")
	fmt.Println("2 - A vista, cartao de credito (5% de desconto)")
	fmt.Println("3 - Em 2 vezes, preco normal sem juros")
	fmt.Println("4 - Em 3 vezes, preco normal + 10% de juros")
	fmt.Print("Digite o codigo da opcao desejada: ")
	fmt.Scan(&codigo)

	switch codigo {
	case 1:
		valorFinal = preco - (preco * 0.10)
	case 2:
		valorFinal = preco - (preco * 0.05)
	case 3:
		valorFinal = preco
	case 4:
		valorFinal = preco + (preco * 0.10)
	default:
		fmt.Println("Codigo invalido!")
		return
	}

	fmt.Printf("Valor a pagar = R$ %.2f\n", valorFinal)
}
