package main

import "fmt"

var n int
var x int

func pot(x int, n int) int {

	if n==0 {

		return 1	

	} else if n==1 {

		return x

	} else {

	return x*pot(x, n - 1)
	
  }

}

func main() {

	fmt.Print("digite o valor de x:\n", x)
	fmt.Scan(&x)

	fmt.Print("digite o valor de n:\n", n)
	fmt.Scan(&n)

    result := pot(x, n)
	fmt.Println("O resultado é:",result)
}