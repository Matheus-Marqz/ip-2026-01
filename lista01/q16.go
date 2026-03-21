// Reajuste salarial
package main

import "fmt"

func main() {
	var salario float64
	fmt.Scanln(&salario)
	if salario <= 300 {
		fmt.Printf("SALARIO COM REAJUSTE = %.2f", 1.5*salario)
	} else {
		fmt.Printf("SALARIO COM REAJUSTE = %.2f", 1.3*salario)
	}
}
