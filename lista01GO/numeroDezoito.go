package main

import "fmt"

func main() {
	var a1, r, n int

	
	fmt.Println("Digite o valor inicial da progressão, a razão e o número de elementos separados por espaço:")
	fmt.Scan(&a1, &r, &n)

	
	soma := calcularSomaProgressaoAritmetica(a1, r, n)

	
	fmt.Println("A soma dos", n, "primeiros termos da progressão aritmética é:", soma)
}

func calcularSomaProgressaoAritmetica(a1, r, n int) int {
	soma := 0
	termo := a1

	
	for i := 0; i < n; i++ {
		soma += termo
		termo += r
	}

	return soma
}
