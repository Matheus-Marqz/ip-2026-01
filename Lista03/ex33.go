package main

import "fmt"

func main() {
	var n1, n2, quociente, resto int

	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)
	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	if n2 <= 0 {
		fmt.Println("Valor invalido")
		return
	}

	resto = n1

	for resto >= n2 {
		resto = resto - n2
		quociente++
	}

	fmt.Println("Quociente:", quociente)
	fmt.Println("Resto:", resto)
}
