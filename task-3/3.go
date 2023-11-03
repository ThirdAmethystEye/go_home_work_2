package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	data := make([]int, n)
	a := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		data[i] = a

	}
	last_elem := data[n-1]
	for i := n - 1; i > 0; i-- {
		data[i] = data[i-1]
	}
	data[0] = last_elem
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", data[i])
	}
}
