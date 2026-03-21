// Cálculo do Determinante de uma Matriz Quadrada de Duas Dimensoes
package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", a*d-b*c)
}
