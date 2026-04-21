package main

import "fmt"

func main() {
	var n, a1, a2, atual int

	fmt.Print("Digite o primeiro termo: ")
	fmt.Scan(&a1)
	fmt.Print("Digite o segundo termo: ")
	fmt.Scan(&a2)
	fmt.Print("Digite a quantidade de termos: ")
	fmt.Scan(&n)

	fmt.Println(a1)
	fmt.Println(a2)

	for i := 3; i <= n; i++ {
		if i%2 != 0 {
			atual = a2 + a1
		} else {
			atual = a2 - a1
		}

		fmt.Println(atual)
		a1 = a2
		a2 = atual
	}
}
