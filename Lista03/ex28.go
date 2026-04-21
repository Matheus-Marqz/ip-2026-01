package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	var sinal float64 = 1
	var denominador float64 = 1

	for i := 1; i <= 51; i++ {
		s = s + sinal*(1/(denominador*denominador*denominador))
		denominador = denominador + 2
		sinal = sinal * -1
	}

	pi := math.Cbrt(s * 32)

	fmt.Println("Pi =", pi)
}
