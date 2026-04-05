package main

import "fmt"

func main() {
	const salarioMinimo = 788.0
	const valorHoraExtra = 10.0

	var matricula int
	var horasExtras float64
	var salarioHoraExtra, salarioBruto float64
	var descontoINSS, descontoIR float64
	var salarioLiquido float64

	fmt.Println("Digite a matricula do funcionario:")
	fmt.Scan(&matricula)

	fmt.Println("Digite a quantidade de horas-extras trabalhadas:")
	fmt.Scan(&horasExtras)

	salarioHoraExtra = horasExtras * valorHoraExtra
	salarioBruto = 3*salarioMinimo + salarioHoraExtra

	if salarioBruto > 1500 {
		descontoINSS = salarioBruto * 0.12
	}

	if salarioBruto > 2000 {
		descontoIR = salarioBruto * 0.20
	}

	salarioLiquido = salarioBruto - descontoINSS - descontoIR

	fmt.Println("\nRESULTADO:")
	fmt.Printf("Matricula: %d\n", matricula)
	fmt.Printf("Salario bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Salario liquido: R$ %.2f\n", salarioLiquido)
}
