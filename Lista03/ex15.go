package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite a quantidade de termos: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(i * i)
	}
}
