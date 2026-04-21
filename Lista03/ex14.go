package main

import "fmt"

func primo(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n1, n2 int

	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)
	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)

	if n1 > n2 {
		n1, n2 = n2, n1
	}

	for i := n1; i <= n2; i++ {
		if primo(i) {
			fmt.Println(i)
		}
	}
}
