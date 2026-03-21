// Conta de Água
package main

import "fmt"

func main() {
	var conta int
	var consumo float64
	var consumidor string
	fmt.Scanln(&conta, &consumo, &consumidor)
	fmt.Printf("CONTA = %d\n", conta)
	switch consumidor {
	case "R":
		fmt.Printf("VALOR DA CONTA = %.2f", 5+0.05*consumo)
	case "C":
		fmt.Printf("VALOR DA CONTA = %.2f", 500+0.25*(consumo-80))
	case "I":
		fmt.Printf("VALOR DA CONTA = %.2f", 800+0.04*(consumo-100))
	default:
		fmt.Println("ERRO")
	}
}
