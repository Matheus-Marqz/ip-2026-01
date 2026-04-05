package main

import "fmt"

func main() {
	var numero int
	var nota1, nota2, nota3, mediaExercicios float64
	var mediaFinal float64
	var conceito string

	fmt.Println("Digite o numero de identificacao do aluno:")
	fmt.Scan(&numero)

	fmt.Println("Digite a primeira nota:")
	fmt.Scan(&nota1)

	fmt.Println("Digite a segunda nota:")
	fmt.Scan(&nota2)

	fmt.Println("Digite a terceira nota:")
	fmt.Scan(&nota3)

	fmt.Println("Digite a media dos exercicios:")
	fmt.Scan(&mediaExercicios)

	mediaFinal = (nota1 + nota2*2 + nota3*3 + mediaExercicios) / 7

	switch {
	case mediaFinal >= 9 && mediaFinal <= 10:
		conceito = "A"
	case mediaFinal >= 7.5:
		conceito = "B"
	case mediaFinal >= 6:
		conceito = "C"
	case mediaFinal >= 4:
		conceito = "D"
	default:
		conceito = "E"
	}

	fmt.Println("\nRESULTADO:")
	fmt.Printf("Numero do aluno: %d\n", numero)
	fmt.Printf("Nota 1: %.2f\n", nota1)
	fmt.Printf("Nota 2: %.2f\n", nota2)
	fmt.Printf("Nota 3: %.2f\n", nota3)
	fmt.Printf("Media dos exercicios e: %.2f\n", mediaExercicios)
	fmt.Printf("Media de aproveitamento e: %.2f\n", mediaFinal)
	fmt.Printf("Conceito: %s\n", conceito)

	if conceito == "A" || conceito == "B" || conceito == "C" {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
