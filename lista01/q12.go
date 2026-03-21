// Locadora de charretes
package main

import "fmt"

func main() {
	var horas int
	fmt.Scanln(&horas)
	if horas >= 3 {
		fmt.Printf("O VALOR A PAGAR E = %.2f\n", float64(int(horas/3)*10+(horas%3)*5))
	}
}
