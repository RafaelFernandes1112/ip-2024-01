package main

import "fmt"

func main() {
	var x, y int

	fmt.Println("Digite dois números inteiros separados por um espaço:")
	fmt.Scan(&x, &y)

	if x%2 == 0 {
		imprimirSequenciaPares(x, y)
	} else {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
		fmt.Println()
	}
}

func imprimirSequenciaPares(x, y int) {
	for i := 0; i < y; i++ {
		fmt.Print(x, " ")
		x += 2
	}
	fmt.Println()
}
