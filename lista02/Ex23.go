package main

import "fmt"

func main() {
	var idade int

	fmt.Println("Digite a idade da pessoa:")
	fmt.Scan(&idade)

	if idade < 16 {
		fmt.Println("Nao-eleitor")
	} else if idade >= 18 && idade <= 65 {
		fmt.Println("Eleitor obrigatorio")
	} else {
		fmt.Println("Eleitor facultativo")
	}
}
