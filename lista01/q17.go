// Série de pares
package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scanln(&n1, &n2)
	if n1%2 == 0 {
		for i := 0; i < n2; i++ {
			fmt.Printf("%d ", n1)
			n1 = n1 + 2
		}
	} else {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
	}
}
