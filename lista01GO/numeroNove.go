package main

import "fmt"

func main() {
	var A, B, C float64

	
	fmt.Println("Digite o coeficiente A:")
	fmt.Scan(&A)

	fmt.Println("Digite o coeficiente B:")
	fmt.Scan(&B)

	fmt.Println("Digite o coeficiente C:")
	fmt.Scan(&C)

	
	delta := B*B - 4*A*C

	fmt.Printf("O VALOR DE DELTA E = %.2f\n", delta)