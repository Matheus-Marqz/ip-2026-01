package main

import "fmt"

func main() {
	var x, s, fatorial float64
	var sinal float64 = -1

	fmt.Print("Digite o valor de X: ")
	fmt.Scan(&x)

	s = x
	fatorial = 1

	for i := 1; i < 20; i++ {
		fatorial = fatorial * float64(i)
		s = s + sinal*(x/fatorial)
		sinal = sinal * -1
	}

	fmt.Println("S =", s)
}
