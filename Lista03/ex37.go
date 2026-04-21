package main

import "fmt"

func main() {
	var numeroOctal string
	var decimal int

	fmt.Print("Digite um numero na base 8: ")
	fmt.Scan(&numeroOctal)

	for i := 0; i < len(numeroOctal); i++ {
		digito := int(numeroOctal[i] - '0')
		decimal = decimal*8 + digito
	}

	fmt.Println("Decimal:", decimal)
}
