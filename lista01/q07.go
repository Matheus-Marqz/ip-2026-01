// Conversões para o Sistema Métrico
package main

import "fmt"

func main() {
	var f, p float64
	fmt.Scanln(&f)
	fmt.Scanln(&p)
	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", 5*(f-32)/9)
	fmt.Printf("O VALOR DE CHUVA E = %.2f", p*25.4)
}
