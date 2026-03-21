// Volume da Pirâmide de Base Hexagonal
package main

import (
	"fmt"
	"math"
)

func main() {
	var altura, aresta float64
	fmt.Scanln(&altura, &aresta)
	ab := 3 * math.Pow(aresta, 2) * math.Sqrt(3) / 2
	v := 1.0 / 3.0 * ab * altura
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", v)
}
