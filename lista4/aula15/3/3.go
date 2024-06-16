package main

import "fmt"

var valor int

func main() {

	vetor := []int{1, 2, 3, 4, 5}

	inv(vetor, 0, len(vetor)-1)

	fmt.Printf("a inversao é:%d", vetor)

}

func inv(array[]int, p int, u int) {
	if p >= u {
		return 
	}
			
	array[p], array[u] = array[u], array[p]

	inv(array, p+1, u-1)
}
	
