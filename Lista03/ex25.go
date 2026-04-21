package main

import "fmt"

func main() {
	var s float64
	var numerador float64 = 1
	var sinal float64 = 1

	for denominador := 15; denominador >= 1; denominador-- {
		s = s + sinal*(numerador/float64(denominador*denominador))
		numerador = numerador * 2
		sinal = sinal * -1
	}

	fmt.Println("S =", s)
}
