package main

import "fmt"

func main() {
	var h float64
	var numerador int = 1

	for denominador := 1; denominador <= 50; denominador++ {
		h = h + float64(numerador)/float64(denominador)
		numerador = numerador + 2
	}

	fmt.Println("H =", h)
}
