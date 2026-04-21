package main

import "fmt"

func main() {
	var n int
	var hexadecimal string
	digitos := "0123456789ABCDEF"

	fmt.Print("Digite um numero na base 10: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Hexadecimal: 0")
		return
	}

	for n > 0 {
		resto := n % 16
		hexadecimal = string(digitos[resto]) + hexadecimal
		n = n / 16
	}

	fmt.Println("Hexadecimal:", hexadecimal)
}
