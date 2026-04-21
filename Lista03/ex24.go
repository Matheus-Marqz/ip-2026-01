package main

import "fmt"

func main() {
	var a, seno float64

	fmt.Println("A\tSen(A)")

	for a = 0.0; a <= 6.3000001; a = a + 0.1 {
		seno = a - (a*a*a)/6 + (a*a*a*a*a)/120 - (a*a*a*a*a*a*a)/5040
		fmt.Printf("%.1f\t%.6f\n", a, seno)
	}
}
