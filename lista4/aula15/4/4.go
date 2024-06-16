package main

import (
	"fmt"
)

func decimalParaBinario(decimal int) string {

	if decimal == 0 {
		return "0"
	}

	digito := decimal % 2

	quociente := decimal / 2
	binario := decimalParaBinario(quociente)

	return fmt.Sprintf("%s%d", binario, digito)
}

func main() {

	var numeroDecimal int

	fmt.Printf("digite o numero decimal:%d\n", numeroDecimal)
	fmt.Scan(&numeroDecimal)

	binario := decimalParaBinario(numeroDecimal)

	fmt.Printf("O número %d em binário é: %s\n", numeroDecimal, binario)
}
