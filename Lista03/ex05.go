package main

import "fmt"

func main() {
	var idade, qtdMais50, qtdEntre10e20, qtdPesoMenor40, totalPessoas, continuar int
	var altura, peso, somaAlturas float64

	for {
		fmt.Print("Digite a idade: ")
		fmt.Scan(&idade)
		fmt.Print("Digite a altura: ")
		fmt.Scan(&altura)
		fmt.Print("Digite o peso: ")
		fmt.Scan(&peso)

		totalPessoas++

		if idade > 50 {
			qtdMais50++
		}

		if idade >= 10 && idade <= 20 {
			somaAlturas = somaAlturas + altura
			qtdEntre10e20++
		}

		if peso < 40 {
			qtdPesoMenor40++
		}

		fmt.Print("Digite 1 para continuar ou outro valor para encerrar: ")
		fmt.Scan(&continuar)

		if continuar != 1 {
			break
		}
	}

	fmt.Println("Quantidade de pessoas com idade superior a 50 anos:", qtdMais50)

	if qtdEntre10e20 > 0 {
		fmt.Println("Media das alturas entre 10 e 20 anos:", somaAlturas/float64(qtdEntre10e20))
	} else {
		fmt.Println("Media das alturas entre 10 e 20 anos: 0")
	}

	fmt.Println("Porcentagem de pessoas com peso inferior a 40 kg:", float64(qtdPesoMenor40)*100/float64(totalPessoas))
}
