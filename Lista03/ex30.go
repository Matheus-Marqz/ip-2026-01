package main

import (
	"fmt"
	"math"
)

func main() {
	var r, volume float64

	fmt.Println("Raio\tVolume")

	for r = 0.0; r <= 20.000001; r = r + 0.5 {
		volume = (4.0 / 3.0) * math.Pi * r * r * r
		fmt.Printf("%.1f\t%.4f\n", r, volume)
	}
}
