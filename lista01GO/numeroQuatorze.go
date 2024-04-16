package main

import (
	"fmt"
	"math"
)

func main() {
	var altura, aresta float64

	fmt.Println("Digite a altura da pirâmide e o comprimento da aresta do hexágono da base (em metros), separados por um espaço:")
	fmt.Scan(&altura, &aresta)

	volume := calcularVolume(altura, aresta)

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", volume)
}

func calcularVolume(altura, aresta float64) float64 {
	areaBase := (3 * math.Pow(aresta, 2) * math.Sqrt(3)) / 2

	volume := (1.0 / 3.0) * areaBase * altura

	return volume
}
