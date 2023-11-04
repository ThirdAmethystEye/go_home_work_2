package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	answer := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "1" {
			answer += "one"
		} else {
			answer += string(s[i])
		}
	}
	println(answer)
}
