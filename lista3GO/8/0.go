package main

import "fmt"

func main() {
	var n int
	fmt.Println("Entrada")
	fmt.Scanln(&n)
	array := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&array[i])
	}
	fmt.Print("Saída\n")
	for _, v := range array {
		fmt.Printf("%b \n", v)
	}

}
