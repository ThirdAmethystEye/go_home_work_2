package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	data := make([]int, n)
	a := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		data[i] = a

	}
	sort.Ints(data)

	prev := 0
	counter := 0
	if len(data) > 0 {
		prev = data[0]
		counter += 1
	}
	for i := 1; i < n; i++ {
		if data[i] != prev {
			prev = data[i]
			counter += 1
		}
	}
	fmt.Printf("%d", counter)
}
