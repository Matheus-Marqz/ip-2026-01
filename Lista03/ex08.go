package main

import "fmt"

func main() {
	var soma int

	for i := 1; i <= 20; i++ {
		fmt.Println(i)
		soma = soma + i
	}

	fmt.Println("Soma:", soma)
}
