package main

import "fmt"

func main() {
	var s float64

	for i := 1; i <= 37; i++ {
		a := 38 - i
		b := 39 - i
		s = s + float64(a*b)/float64(i)
	}

	fmt.Println("S =", s)
}
