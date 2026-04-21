package main

import "fmt"

func main() {
	var numero, numeroGordo, numeroMagro int
	var peso, pesoGordo, pesoMagro float64

	for i := 1; i <= 90; i++ {
		fmt.Println("Boi", i)
		fmt.Print("Digite o numero de identificacao: ")
		fmt.Scan(&numero)
		fmt.Print("Digite o peso: ")
		fmt.Scan(&peso)

		if i == 1 {
			numeroGordo = numero
			numeroMagro = numero
			pesoGordo = peso
			pesoMagro = peso
		} else {
			if peso > pesoGordo {
				pesoGordo = peso
				numeroGordo = numero
			}
			if peso < pesoMagro {
				pesoMagro = peso
				numeroMagro = numero
			}
		}
	}

	fmt.Println("Boi mais gordo:", numeroGordo, "Peso:", pesoGordo)
	fmt.Println("Boi mais magro:", numeroMagro, "Peso:", pesoMagro)
}
