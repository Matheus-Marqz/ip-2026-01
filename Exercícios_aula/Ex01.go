package main

import "fmt"

func main() {
	var l1, l2, l3 float64
	fmt.Scan(&l1, &l2, &l3)
	if l1+l2 > l3 && l2+l3 > l1 && l1+l3 > l2 {
		if l1 == l2 && l1 == l3 {
			fmt.Println("Equilátero")
		} else if l1 != l2 && l2 != l3 && l1 != l3 {
			fmt.Println("Escaleno")
		} else {
			fmt.Println("Isósceles")
		}

	} else {
	    fmt.Println("Não é um triângulo")
	}
}
