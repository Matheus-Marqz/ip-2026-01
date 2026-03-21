// Consumo de energia
package main

import "fmt"

func main() {
	var sm, kw float64
	fmt.Scanln(&sm)
	fmt.Scanln(&kw)
	ckw := sm * (0.7 / 100)
	fmt.Printf("Custo por KW: R$ %.2f\n", ckw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", ckw*kw)
	fmt.Printf("Custo com desconto: R$ %.2f\n", 0.9*kw*ckw)
}
