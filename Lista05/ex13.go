package main

import "fmt"

func main() {
	const max = 100
	var codigo, meses int

	cod1, cod2, cod3 := 0, 0, 0
	mes1, mes2, mes3 := 1<<30, 1<<30, 1<<30
	qtd := 0

	fmt.Println("Digite o codigo do empregado e os meses de trabalho.")
	fmt.Println("Para encerrar, digite 0 0.")

	for qtd < max {
		fmt.Scan(&codigo, &meses)

		if codigo == 0 && meses == 0 {
			break
		}

		qtd++

		if meses < mes1 {
			mes3, cod3 = mes2, cod2
			mes2, cod2 = mes1, cod1
			mes1, cod1 = meses, codigo
		} else if meses < mes2 {
			mes3, cod3 = mes2, cod2
			mes2, cod2 = meses, codigo
		} else if meses < mes3 {
			mes3, cod3 = meses, codigo
		}
	}

	if qtd == 0 {
		fmt.Println("Nenhum empregado informado.")
		return
	}

	fmt.Println("Tres empregados mais recentes:")
	fmt.Printf("1º - Codigo: %d | Meses de trabalho: %d\n", cod1, mes1)
	if qtd >= 2 {
		fmt.Printf("2º - Codigo: %d | Meses de trabalho: %d\n", cod2, mes2)
	}
	if qtd >= 3 {
		fmt.Printf("3º - Codigo: %d | Meses de trabalho: %d\n", cod3, mes3)
	}
}
