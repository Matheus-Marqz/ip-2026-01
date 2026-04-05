package main

import "fmt"

func main() {
	var conta int
	var tipo string
	var consumo int
	var valor float64
    fmt.Println("Escreva a conta do cliente, o tipo e e o consumo em mestros cúbicos(EX: 123 residencial 32")
	fmt.Scan(&conta, &tipo, &consumo)

	switch tipo {
	case "residencial":
		valor = 5 + float64(consumo)*0.05

	case "comercial":
		if consumo <= 80 {
			valor = 500
		} else {
			valor = 500 + float64(consumo-80)*0.25
		}

	case "industrial":
		if consumo <= 100 {
			valor = 800
		} else {
			valor = 800 + float64(consumo-100)*0.04
		}
	}

	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR A PAGAR = %.2f\n", valor)
}
