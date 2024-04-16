package main

import (
	"fmt"
	"math"
)

func main() {
	var N int

	fmt.Println("Digite um valor inteiro N (5 < N < 2000):")
	fmt.Scan(&N)

	if N <= 5 || N >= 2000 {
		fmt.Println("N fora do intervalo permitido.")
		return
	}

	for i := 2; i <= N; i += 2 {
		quadrado := int(math.Pow(float64(i), 2))
		fmt.Printf("%d^2 = %d\n", i, quadrado)
	}
}