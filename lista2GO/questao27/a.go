package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactorization(n int) string {
	if n <= 1 {
		return fmt.Sprintf("Fatoracao nao e possivel para o numero %d!\n", n)
	}

	factors := ""
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			if factors != "" {
				factors += " x "
			}
			factors += fmt.Sprintf("%d", i)
			n /= i
		}
	}
	return factors + "\n"
}

func main() {
	var n int
	for {
		fmt.Scan(&n)
		if n > 1 {
			fmt.Printf("%d = %s", n, primeFactorization(n))
		} else {
			fmt.Printf(primeFactorization(n))
		}
	}
}
