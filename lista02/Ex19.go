package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var r, h float64
	var volume, area, g float64
    fmt.Print("Qual figura você deseja? (1- cone; 2- cilindro; 3- esfera) ")
	fmt.Scan(&opcao)

	switch opcao {
	case 1: // cone
	    fmt.Print("Digite o raio e a altura do cone: ")
		fmt.Scan(&r, &h)
		g = math.Sqrt(r*r + h*h)
		volume = (math.Pi * r * r * h) / 3.0
		area = math.Pi * r * (r + g)

	case 2: // cilindro
	    fmt.Print("Digite o raio e a altura do cilindro: ")
		fmt.Scan(&r, &h)
		volume = math.Pi * r * r * h
		area = 2 * math.Pi * r * (r + h)

	case 3: // esfera
	    fmt.Print("Digite o raio da esfera: ")
		fmt.Scan(&r)
		volume = (4.0 * math.Pi * r * r * r) / 3.0
		area = 4 * math.Pi * r * r

	default:
		fmt.Println("Opcao invalida!")
		return
	}

	fmt.Printf("VOLUME = %.2f\n", volume)
	fmt.Printf("AREA = %.2f\n", area)
}
