package main

import "fmt"

func main() {
	var v [10]int
	somaPares := 0
	qtdImpares := 0

	fmt.Println("Digite 10 numeros inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Print("Numeros pares digitados: ")
	for i := 0; i < 10; i++ {
		if v[i]%2 == 0 {
			fmt.Printf("%d ", v[i])
			somaPares += v[i]
		}
	}
	fmt.Println()

	fmt.Printf("Soma dos numeros pares: %d\n", somaPares)

	fmt.Print("Numeros impares digitados: ")
	for i := 0; i < 10; i++ {
		if v[i]%2 != 0 {
			fmt.Printf("%d ", v[i])
			qtdImpares++
		}
	}
	fmt.Println()

	fmt.Printf("Quantidade de numeros impares: %d\n", qtdImpares)
}
