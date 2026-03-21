// Arrecadação de Jogos
package main

import "fmt"

func main() {
	var t int
	var np, g, a, c, p float64
	var l []float64
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&np, &p, &g, &a, &c)
		l = append(l, 0.01*((np*p)*1+(np*g)*5+(np*a)*10+(np*c)*20))
	}
	for i := 0; i < t; i++ {
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i+1, l[i])
	}
}
