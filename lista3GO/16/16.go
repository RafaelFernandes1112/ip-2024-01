package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	var chegadas []int
	for i := 0; i < N; i++ {
		var tempoChegada int
		fmt.Scan(&tempoChegada)
		chegadas = append(chegadas, tempoChegada)
	}

	alunosPresentes := 0
	for _, tempo := range chegadas {
		if tempo <= 0 {
			alunosPresentes++
		}
	}

	if alunosPresentes < K {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
		for i := N - 1; i >= 0; i-- {
			if chegadas[i] <= 0 {
				fmt.Print(i+1, " ")
			}
		}
		fmt.Println()
	}
}