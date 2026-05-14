package main

import "fmt"

func main() {
	var v1 [10]int
	var v2 [10]int
	var intercalado [20]int

	fmt.Println("Digite os 10 elementos do primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v1[i])
	}

	fmt.Println("Digite os 10 elementos do segundo vetor:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v2[i])
	}

	k := 0
	for i := 0; i < 10; i++ {
		intercalado[k] = v1[i]
		k++
		intercalado[k] = v2[i]
		k++
	}

	fmt.Println("Vetor resultante da intercalacao:")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", intercalado[i])
	}
	fmt.Println()
}
