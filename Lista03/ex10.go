package main

import "fmt"

func main() {
	var n int
	var a, b, c int

	fmt.Print("Digite a quantidade de termos: ")
	fmt.Scan(&n)

	a = 0
	b = 1

	fmt.Println(a)
	fmt.Println(b)

	for i := 3; i <= n; i++ {
		c = a + b
		fmt.Println(c)
		a = b
		b = c
	}
}
