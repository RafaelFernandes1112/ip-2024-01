package main

import "fmt"

func main() {
	var conta int
	var consumo, valorConta float64
	var tipoConsumidor rune

	fmt.Println("Digite a conta do cliente, o consumo de água (em metros cúbicos) e o tipo de consumidor (R - Residencial, C - Comercial, I - Industrial), separados por espaço:")
	fmt.Scan(&conta, &consumo, &tipoConsumidor)

	switch tipoConsumidor {
	case 'R':
		valorConta = 5.00 + 0.05*consumo
	case 'C':
		if consumo <= 80 {
			valorConta = 500.00
		} else {
			valorConta = 500.00 + 0.25*(consumo-80)
		}
	case 'I':
		if consumo <= 100 {
			valorConta = 800.00
		} else {
			valorConta = 800.00 + 0.04*(consumo-100)
		}
	}

	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valorConta)
}
