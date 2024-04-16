package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Digite o salário do funcionário:")
	fmt.Scan(&salario)

	var salarioReajustado float64
	if salario <= 300.00 {
		salarioReajustado = salario * 1.5 // Reajuste de 50%
	} else {
		salarioReajustado = salario * 1.3 // Reajuste de 30%
	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salarioReajustado)
}
