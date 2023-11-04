package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	data := make([][]int, n)
	for i := 0; i < n; i++ {
		data[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a := 0
			if j != 0 {
				fmt.Scanf("%d", &a)
			} else {
				fmt.Scanf("\n%d", &a)
			}
			data[i][j] = a
		}
	}

	flag := true
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if data[i][j] != data[j][i] {
				flag = false
				break
			}
		}
	}
	if flag == true {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
