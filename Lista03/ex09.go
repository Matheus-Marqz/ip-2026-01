package main

import "fmt"

func main() {
	var n, aprovados, exame, reprovados int
	var nota1, nota2, media, somaMedias float64

	fmt.Print("Digite a quantidade de alunos: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println("Aluno", i)
		fmt.Print("Digite a primeira nota: ")
		fmt.Scan(&nota1)
		fmt.Print("Digite a segunda nota: ")
		fmt.Scan(&nota2)

		media = (nota1 + nota2) / 2
		somaMedias = somaMedias + media

		fmt.Println("Media:", media)

		if media < 3 {
			fmt.Println("Reprovado")
			reprovados++
		} else if media < 7 {
			fmt.Println("Exame")
			exame++
		} else {
			fmt.Println("Aprovado")
			aprovados++
		}
	}

	fmt.Println("Total de aprovados:", aprovados)
	fmt.Println("Total de exame:", exame)
	fmt.Println("Total de reprovados:", reprovados)
	fmt.Println("Media da classe:", somaMedias/float64(n))
}
