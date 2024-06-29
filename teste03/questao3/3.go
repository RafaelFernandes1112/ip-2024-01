package main

import "fmt"

func BubbleSort(l []int) {
	n := len(l)
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if l[j-1] > l[j] {
				l[j-1], l[j] = l[j], l[j-1]
			}
		}
	}
}
func main() {
	nums := []int{23, 45, 56, 11, 10, 8, 78}
	fmt.Println("Lista original:", nums)
	BubbleSort(nums)
	fmt.Println("Lista ordenada:", nums)
}
