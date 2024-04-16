package main

import "fmt"

func main() {
	var horas, minutos, segundos int

	fmt.Println("Digite o valor em horas:")
	fmt.Scan(&horas)
	fmt.Println("Digite o valor em minutos:")
	fmt.Scan(&minutos)
	fmt.Println("Digite o valor em segundos:")
	fmt.Scan(&segundos)

	totalSegundos := converterParaSegundos(horas, minutos, segundos)

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", totalSegundos)
}

func converterParaSegundos(horas, minutos, segundos int) int {

	totalSegundos := horas*3600 + minutos*60 + segundos
	return totalSegundos
}
