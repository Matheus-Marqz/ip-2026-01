// Custo da Lata de Cerveja
package main

import "fmt"

func main() {
	var r, h float64
	const pi float64 = 3.14159
	fmt.Scanln(&r)
	fmt.Scanln(&h)
	fmt.Printf("O VALOR DO CUSTO E = %.2f", (2*(pi*r*r)+(2*pi*r*h))*100)
}
