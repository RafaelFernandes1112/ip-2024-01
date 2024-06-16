package main

import "fmt"

func main() {
	vetor := []int{1, 2, 3, 14, 15}


	resultado := soma(vetor, 0)

	fmt.Printf("A soma dos valores é: %d", resultado)
}

func soma(array []int, v int) int {

	if v == len(array) {
		return 0
	}

	return array[v] + soma(array, v+1)
}
