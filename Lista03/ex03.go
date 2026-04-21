package main

import "fmt"

func main() {
	var salarioCarlos, salarioJoao float64
	var meses int

	fmt.Print("Digite o salario de Carlos: ")
	fmt.Scan(&salarioCarlos)

	salarioJoao = salarioCarlos / 3

	for salarioJoao < salarioCarlos {
		salarioCarlos = salarioCarlos * 1.02
		salarioJoao = salarioJoao * 1.05
		meses++
	}

	fmt.Println("Meses necessarios:", meses)
}
