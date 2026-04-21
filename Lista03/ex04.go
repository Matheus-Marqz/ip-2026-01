package main

import "fmt"

func main() {
	var n, i int
	var quadradoPerfeito bool

	for {
		fmt.Print("Digite um numero inteiro (<= 0 para encerrar): ")
		fmt.Scan(&n)

		if n <= 0 {
			break
		}

		quadradoPerfeito = false
		i = 1

		for i*i <= n {
			if i*i == n {
				quadradoPerfeito = true
				break
			}
			i++
		}

		if quadradoPerfeito {
			fmt.Println("E quadrado perfeito")
		} else {
			fmt.Println("Nao e quadrado perfeito")
		}
	}
}
