package main

import (
	"fmt"
)

func main() {
	var a int
	i := 0
	fmt.Scan(&a)
	preAns := make([]int, 0, 0)
	for a > 0 {
		preAns = append(preAns, (a%10)*(a%10))
		a /= 10
		i += 1
	}
	ans := 0
	for i-1 >= 0 {
		q := preAns[i-1]
		if q > 10 {
			ans = ans*100 + q
		} else {
			ans = ans*10 + q
		}
		i -= 1
	}
	fmt.Print(ans)
}





