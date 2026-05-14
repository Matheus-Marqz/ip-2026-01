package main

import "fmt"

func main() {
	var v [10]int
	encontrou := false

	fmt.Println("Digite 10 numeros inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Println("Elementos repetidos e quantidade de repeticoes:")
	for i := 0; i < 10; i++ {
		jaMostrado := false

		for j := 0; j < i; j++ {
			if v[i] == v[j] {
				jaMostrado = true
				break
			}
		}

		if jaMostrado {
			continue
		}

		cont := 0
		for j := 0; j < 10; j++ {
			if v[i] == v[j] {
				cont++
			}
		}

		if cont > 1 {
			fmt.Printf("Numero %d aparece %d vezes\n", v[i], cont)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Nao ha elementos repetidos.")
	}
}
