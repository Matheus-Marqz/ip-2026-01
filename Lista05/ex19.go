package main

import "fmt"

func main() {
	var num [10]int
	var divis [5]int

	fmt.Println("Digite os 10 elementos do primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&num[i])
	}

	fmt.Println("Digite os 5 elementos do segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&divis[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Numero %d:\n", num[i])
		encontrou := false

		for j := 0; j < 5; j++ {
			if divis[j] != 0 && num[i]%divis[j] == 0 {
				fmt.Printf("Divisivel por %d na posicao %d\n", divis[j], j)
				encontrou = true
			}
		}

		if !encontrou {
			fmt.Println("Nao possui divisores no segundo vetor.")
		}
	}
}
