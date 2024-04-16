package main

import (
	"fmt"
)

func main() {
	
	notas := []float64{3, 4, 5}

	
	var soma float64
	for _, nota := range notas {
		soma += nota
	}
	media := soma / float64(len(notas))

	
	var status string
	if media >= 6 {
		status = "APROVADO"
	} else {
		status = "REPROVADO"
	}

	
	fmt.Printf("MEDIA = %.2f\n%s\n", media, status)
}