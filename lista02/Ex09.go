package main

import "fmt"

func main() {
	var compra, venda float64

	fmt.Println("Digite o valor da compra:")
	fmt.Scan(&compra)

	if compra < 0 {
		fmt.Println("Valor de compra invalido!")
		return
	}

	if compra < 10 {
		venda = compra + compra*0.70
	} else if compra < 30 {
		venda = compra + compra*0.50
	} else if compra < 50 {
		venda = compra + compra*0.40
	} else {
		venda = compra + compra*0.30
	}

	fmt.Printf("Valor da venda = R$ %.2f\n", venda)
}
