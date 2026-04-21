package main

import "fmt"

func main() {
	var n, soma, quantidade, maior, menor, somaPares, qtdPares, qtdImpares int
	var media, mediaPares, percentualImpares float64
	var primeiro bool = true

	for {
		fmt.Print("Digite um numero (30000 para encerrar): ")
		fmt.Scan(&n)

		if n == 30000 {
			break
		}

		soma = soma + n
		quantidade++

		if primeiro {
			maior = n
			menor = n
			primeiro = false
		} else {
			if n > maior {
				maior = n
			}
			if n < menor {
				menor = n
			}
		}

		if n%2 == 0 {
			somaPares = somaPares + n
			qtdPares++
		} else {
			qtdImpares++
		}
	}

	if quantidade > 0 {
		media = float64(soma) / float64(quantidade)
		percentualImpares = float64(qtdImpares) * 100 / float64(quantidade)
	}

	if qtdPares > 0 {
		mediaPares = float64(somaPares) / float64(qtdPares)
	}

	fmt.Println("Soma:", soma)
	fmt.Println("Quantidade:", quantidade)
	fmt.Println("Media:", media)

	if quantidade > 0 {
		fmt.Println("Maior numero:", maior)
		fmt.Println("Menor numero:", menor)
	}

	fmt.Println("Media dos pares:", mediaPares)
	fmt.Println("Percentual de impares:", percentualImpares)
}
