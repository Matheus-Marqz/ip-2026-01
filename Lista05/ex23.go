package main

import "fmt"

func mostrarDisponiveis(vetor [24]int, tipo string) {
	fmt.Printf("Poltronas disponiveis na %s: ", tipo)
	encontrou := false

	for i := 0; i < 24; i++ {
		if vetor[i] == 0 {
			fmt.Printf("%d ", i+1)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Print("nenhuma")
	}
	fmt.Println()
}

func lotado(vetor [24]int) bool {
	for i := 0; i < 24; i++ {
		if vetor[i] == 0 {
			return false
		}
	}
	return true
}

func main() {
	var janela [24]int
	var corredor [24]int
	var opcao, poltrona int

	for {
		if lotado(janela) && lotado(corredor) {
			fmt.Println("Onibus completamente cheio.")
			return
		}

		fmt.Println("\nMenu")
		fmt.Println("1. Comprar poltrona na janela")
		fmt.Println("2. Comprar poltrona no corredor")
		fmt.Println("3. Mostrar situacao das poltronas")
		fmt.Println("4. Encerrar")
		fmt.Print("Opcao: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			if lotado(janela) {
				fmt.Println("Nao existem poltronas livres na janela.")
				continue
			}

			mostrarDisponiveis(janela, "janela")
			fmt.Print("Escolha o numero da poltrona: ")
			fmt.Scan(&poltrona)

			if poltrona < 1 || poltrona > 24 {
				fmt.Println("Poltrona invalida.")
			} else if janela[poltrona-1] == 1 {
				fmt.Println("Poltrona ja ocupada.")
			} else {
				janela[poltrona-1] = 1
				fmt.Println("Venda realizada com sucesso.")
			}

		case 2:
			if lotado(corredor) {
				fmt.Println("Nao existem poltronas livres no corredor.")
				continue
			}

			mostrarDisponiveis(corredor, "corredor")
			fmt.Print("Escolha o numero da poltrona: ")
			fmt.Scan(&poltrona)

			if poltrona < 1 || poltrona > 24 {
				fmt.Println("Poltrona invalida.")
			} else if corredor[poltrona-1] == 1 {
				fmt.Println("Poltrona ja ocupada.")
			} else {
				corredor[poltrona-1] = 1
				fmt.Println("Venda realizada com sucesso.")
			}

		case 3:
			fmt.Println("Situacao da janela:", janela)
			fmt.Println("Situacao do corredor:", corredor)

		case 4:
			fmt.Println("Programa encerrado.")
			return

		default:
			fmt.Println("Opcao invalida.")
		}
	}
}
