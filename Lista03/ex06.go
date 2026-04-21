package main

import "fmt"

func main() {
	var n, i int
	var triangular bool

	fmt.Print("Digite um numero inteiro positivo: ")
	fmt.Scan(&n)

	for i = 1; i*(i+1)*(i+2) <= n; i++ {
		if i*(i+1)*(i+2) == n {
			triangular = true
			break
		}
	}

	if triangular {
		fmt.Println("E triangular")
	} else {
		fmt.Println("Nao e triangular")
	}
}
