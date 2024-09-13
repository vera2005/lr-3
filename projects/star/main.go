package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var a string
	fmt.Scan(&a)
	qwe := make([]string, 0, 0)
	srOfRune := []rune(a)
	for i := range srOfRune {
		if i != utf8.RuneCountInString(a)-1 {
			qwe = append(qwe, string(srOfRune[i]), "*")
		} else {
			qwe = append(qwe, string(srOfRune[i]))
		}
	}
	ans := strings.Join(qwe, "")
	fmt.Print(ans)

}
