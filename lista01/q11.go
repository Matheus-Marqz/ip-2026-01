// Divisível por 3 e 5
package main

import "fmt"

func main() {
	var v int
	fmt.Scanln(&v)
	if (v%5 == 0) && (v%3 == 0) {
		fmt.Println("O NUMERO E DIVISIVEL")
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL")
	}
}
