// Composição Inteira (+)
package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)
	if n1 > 9 || n2 > 9 || n3 > 9 {
		fmt.Println("DIGITO INVALIDO")
	} else {
		j := n1*100 + n2*10 + n3
		fmt.Printf("%d, ", j)
		fmt.Println(j * j)
	}
}
