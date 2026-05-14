package main

import "fmt"

func buscarConta(codigos [10]int, codigo int) int {
	for i := 0; i < 10; i++ {
		if codigos[i] == codigo {
			return i
		}
	}
	return -1
}

func codigoExisteAte(codigos [10]int, codigo int, limite int) bool {
	for i := 0; i < limite; i++ {
		if codigos[i] == codigo {
			return true
		}
	}
	return false
}

func main() {
	var codigos [10]int
	var saldos [10]float64
	var opcao int

	fmt.Println("Cadastro de 10 contas bancarias:")

	for i := 0; i < 10; i++ {
		for {
			fmt.Printf("Digite o codigo da conta %d: ", i+1)
			fmt.Scan(&codigos[i])

			if !codigoExisteAte(codigos, codigos[i], i) {
				break
			}

			fmt.Println("Codigo ja cadastrado. Digite outro.")
		}

		fmt.Printf("Digite o saldo da conta %d: ", i+1)
		fmt.Scan(&saldos[i])
	}

	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Efetuar deposito")
		fmt.Println("2. Efetuar saque")
		fmt.Println("3. Consultar o ativo bancario")
		fmt.Println("4. Finalizar o programa")
		fmt.Print("Opcao: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var codigo int
			var valor float64

			fmt.Print("Digite o codigo da conta: ")
			fmt.Scan(&codigo)

			pos := buscarConta(codigos, codigo)
			if pos == -1 {
				fmt.Println("Conta nao encontrada.")
				continue
			}

			fmt.Print("Digite o valor do deposito: ")
			fmt.Scan(&valor)
			saldos[pos] += valor
			fmt.Printf("Deposito realizado. Novo saldo: %.2f\n", saldos[pos])

		case 2:
			var codigo int
			var valor float64

			fmt.Print("Digite o codigo da conta: ")
			fmt.Scan(&codigo)

			pos := buscarConta(codigos, codigo)
			if pos == -1 {
				fmt.Println("Conta nao encontrada.")
				continue
			}

			fmt.Print("Digite o valor do saque: ")
			fmt.Scan(&valor)

			if saldos[pos] >= valor {
				saldos[pos] -= valor
				fmt.Printf("Saque realizado. Novo saldo: %.2f\n", saldos[pos])
			} else {
				fmt.Println("Saldo insuficiente.")
			}

		case 3:
			ativo := 0.0
			for i := 0; i < 10; i++ {
				ativo += saldos[i]
			}
			fmt.Printf("Ativo bancario: %.2f\n", ativo)

		case 4:
			fmt.Println("Programa finalizado.")
			return

		default:
			fmt.Println("Opcao invalida.")
		}
	}
}
