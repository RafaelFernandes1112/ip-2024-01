package main

import (
	"fmt"
)

func main() {
	var n int


	fmt.Println("Digite um número inteiro positivo maior que 1:")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}


	soma := calcularSoma(n)

	fmt.Printf("%.6f\n", soma)
}

func calcularSoma(n int) float64 {
	var soma float64


	for k := 1; k <= n; k++ {
		soma += 1.0 / float64(k)
	}

	return soma
}
