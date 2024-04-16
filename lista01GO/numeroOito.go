package main

import (
	"fmt"
	"math"
)

func main() {
	var raio, altura float64

	fmt.Println("Digite o raio da lata em metros:")
	fmt.Scan(&raio)

	fmt.Println("Digite a altura da lata em metros:")
	fmt.Scan(&altura)

	areaCirculo := math.Pi * raio * raio

	areaLateral := 2 * math.Pi * raio * altura

	areaTotal := 2*areaCirculo + areaLateral

	custo := 100.00 * areaTotal

	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}
