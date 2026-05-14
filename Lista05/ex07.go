package main

import "fmt"

func main() {
	var v [100]int
	numero := 1

	for i := 0; i < 100; i++ {
		v[i] = numero
		numero += 2
	}

	fmt.Println("Os 100 primeiros numeros impares:")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", v[i])
	}
	fmt.Println()
}
