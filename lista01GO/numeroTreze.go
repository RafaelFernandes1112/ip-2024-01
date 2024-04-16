package main

import "fmt"

func main() {
	var nota float64

	fmt.Println("Digite a nota entre 0 e 10 com uma casa decimal:")
	fmt.Scan(&nota)

	conceito := converterParaConceito(nota)

	fmt.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
}

func converterParaConceito(nota float64) string {
	var conceito string

	switch {
	case nota >= 9.0 && nota <= 10:
		conceito = "A"
	case nota >= 7.5 && nota < 9.0:
		conceito = "B"
	case nota >= 6.0 && nota < 7.5:
		conceito = "C"
	default:
		conceito = "D"
	}

	return conceito
}
