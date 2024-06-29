package main

import (
	"fmt"
)

func main() {
	var numCasos int
	fmt.Scanln(&numCasos)

	for caso := 0; caso < numCasos; caso++ {
		var tamanho int
		fmt.Scanln(&tamanho)

		vetor := make([]int, tamanho)
		for i := 0; i < tamanho; i++ {
			fmt.Scan(&vetor[i])
		}

		trocaOpostosSeMenor(vetor, tamanho)

		
		for i := 0; i < tamanho; i++ {
			fmt.Print(vetor[i])
			if i < tamanho-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println() 
	}
}

func troca(vetor []int, x, y int) {
	vetor[x], vetor[y] = vetor[y], vetor[x]
}


func trocaOpostosSeMenor(vetor []int, tamanho int) {
	for i := 0; i < tamanho/2; i++ {
		oposto := tamanho - 1 - i
		if i < oposto && vetor[i] < vetor[oposto] {
			troca(vetor, i, oposto)
		}
	}
}