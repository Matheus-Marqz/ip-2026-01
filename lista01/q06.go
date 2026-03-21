// Conversão de temperatura
package main

import "fmt"

func main() {
	var n int
	var l []float64
	var t float64
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&t)
		l = append(l, t)
	}
	for r := 0; r < n; r++ {
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", l[r], 5*(l[r]-32)/9)
	}
}
