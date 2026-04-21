package main

import "fmt"

func main() {
	var s float64
	var fatorial float64 = 1

	for i := 0; i < 20; i++ {
		if i > 0 {
			fatorial = fatorial * float64(i)
		}
		s = s + float64(100-i)/fatorial
	}

	fmt.Println("S =", s)
}
