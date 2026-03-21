// Conversão de Nota em Conceito
package main

import "fmt"

func main() {
	var nota float64
	var conceito string
	fmt.Scanln(&nota)
	switch {
	case nota >= 9:
		conceito = "A"
	case nota < 9 && nota >= 7.5:
		conceito = "B"
	case nota < 7.5 && nota >= 6:
		conceito = "C"
	case nota < 6:
		conceito = "D"
	}
	fmt.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
}
