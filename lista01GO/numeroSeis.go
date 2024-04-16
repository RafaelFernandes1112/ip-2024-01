package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite o número de temperaturas em Fahrenheit a serem convertidas:")
	fmt.Scan(&n)

	fmt.Println("Digite as temperaturas em Fahrenheit, uma por linha:")
	for i := 0; i < n; i++ {
		var fahrenheit float64
		fmt.Scan(&fahrenheit)

		celsius := (5 * (fahrenheit - 32)) / 9

		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", fahrenheit, celsius)
	}
}
