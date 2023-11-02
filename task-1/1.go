package main

import (
	"fmt"
	"math"
)

func main() {
	a := 0.0
	b := 0.0
	fmt.Scanf("%f\n%f", &a, &b)
	fmt.Println(math.Sqrt(a*a + b*b))
}
