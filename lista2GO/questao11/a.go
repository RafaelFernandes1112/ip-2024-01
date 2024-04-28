package main

import (
	"fmt"
	"math"
)

func main() {
	var n, e float64
	fmt.Scan(&n, &e)

	r := 1.0
	for {
		rNext := (r + n/r) / 2
		error := math.Abs(n - rNext*rNext)
		fmt.Printf("r: %.9f, erro: %.9f\n", rNext, error)
		if error <= e {
			break
		}
		r = rNext
	}
}