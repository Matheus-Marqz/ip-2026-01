package main

import "fmt"

func main() {
	var n int
	var binario string

	fmt.Print("Digite um numero na base 10: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Binario: 0")
		return
	}

	for n > 0 {
		resto := n % 2
		binario = fmt.Sprintf("%d%s", resto, binario)
		n = n / 2
	}

	fmt.Println("Binario:", binario)
}
