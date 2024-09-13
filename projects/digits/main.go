package main

import "fmt"

func main() {
	ans := 0
	var a string
	fmt.Scan(&a)
	srezRune := []rune(a)
	for _, el := range srezRune {
		if int(el) > ans {
			ans = int(el)
		}
	}
	fmt.Print(string(ans))
}
