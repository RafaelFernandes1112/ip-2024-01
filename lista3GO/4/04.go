package main

import "fmt"

func main(){

var n int
fmt.Print("Entrada")
fmt.Scan(&n)

si := make([]int, n, n)
for i := range si {
	fmt.Scan(&si[i])

}

fmt.Print("Saída")
fmt.Printf("%v", si)

}