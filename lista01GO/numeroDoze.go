package main

import "fmt"

func main() {
	var horas int

	fmt.Println("Digite o número de horas que a charrete foi utilizada:")
	fmt.Scan(&horas)

	valor := calcularValor(horas)

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}

func calcularValor(horas int) float64 {
	const taxaPor3Horas = 10.00
	const taxaPorHoraAbaixoDe3 = 5.00

	blocosDe3Horas := horas / 3

	horasRestantes := horas % 3

	valor := float64(blocosDe3Horas*taxaPor3Horas + horasRestantes*taxaPorHoraAbaixoDe3)

	return valor
}
