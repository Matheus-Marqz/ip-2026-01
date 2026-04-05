package main

import "fmt"

func main() {
	var destino, retorno int
	var preco float64

	fmt.Println("Digite o destino:")
	fmt.Println("1 - Regiao Norte")
	fmt.Println("2 - Regiao Nordeste")
	fmt.Println("3 - Regiao Centro-Oeste")
	fmt.Println("4 - Regiao Sul")
	fmt.Scan(&destino)

	fmt.Println("A viagem inclui retorno?")
	fmt.Println("1 - Sim (ida e volta)")
	fmt.Println("2 - Nao (so ida)")
	fmt.Scan(&retorno)

	if retorno != 1 && retorno != 2 {
		fmt.Println("Opcao de retorno invalida!")
		return
	}

	switch destino {
	case 1:
		if retorno == 1 {
			preco = 900
		} else {
			preco = 500
		}
	case 2:
		if retorno == 1 {
			preco = 650
		} else {
			preco = 350
		}
	case 3:
		if retorno == 1 {
			preco = 600
		} else {
			preco = 350
		}
	case 4:
		if retorno == 1 {
			preco = 550
		} else {
			preco = 300
		}
	default:
		fmt.Println("Destino invalido!")
		return
	}

	fmt.Printf("Preco da passagem = R$ %.2f\n", preco)
}
