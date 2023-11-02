package main

import "fmt"

func main() {
	a := 0
	b := 0
	c := 0
	fmt.Scanf("%d\n%d\n%d", &a, &b, &c)
	if (a+b > c) && (b+c > a) && (a+c > b) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
