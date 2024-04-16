package main

import (
	"fmt"
)

func calcularRendaTotal(pessoas int, percentuais []float64) float64 {

	valorIngresso := []float64{1.00, 5.00, 10.00, 20.00}

	var ingressosPorCategoria []float64
	for i := 0; i < len(percentuais); i++ {
		ingressosPorCategoria = append(ingressosPorCategoria, float64(pessoas)*percentuais[i]/100)
	}

	var rendaTotal float64
	for i := 0; i < len(ingressosPorCategoria); i++ {
		rendaTotal += ingressosPorCategoria[i] * valorIngresso[i]
	}

	return rendaTotal
}

func main() {
	var casosTeste int
	fmt.Scanln(&casosTeste)

	for i := 1; i <= casosTeste; i++ {
		var pessoas int
		var percentuais []float64

		fmt.Scanln(&pessoas)
		for j := 0; j < 4; j++ {
			var percentual float64
			fmt.Scan(&percentual)
			percentuais = append(percentuais, percentual)
		}

		rendaTotal := calcularRendaTotal(pessoas, percentuais)
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, rendaTotal)
	}
}
