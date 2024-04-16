package main

import "fmt"

func main() {
	var fahrenheit, polegadas float64

	fmt.Println("Digite a temperatura em Fahrenheit:")
	fmt.Scan(&fahrenheit)

	fmt.Println("Digite a quantidade de chuva em polegadas:")
	fmt.Scan(&polegadas)

	celsius := (5*fahrenheit - 160) / 9

	milimetros := polegadas * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", celsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", milimetros)
}
