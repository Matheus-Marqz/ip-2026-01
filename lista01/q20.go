// Tempo em segundos
package main

import "fmt"

func main() {
	var h, m, s int
	fmt.Scanln(&h)
	fmt.Scanln(&m)
	fmt.Scanln(&s)
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", s+m*60+h*60*60)
}
