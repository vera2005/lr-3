package main

import (
	"fmt"
	"math"
)

var k, p, v float64

func main() {
	fmt.Scan(&k)
	fmt.Scan(&p)
	fmt.Scan(&v)
	fmt.Print(T())
}

func T() float64 {
	return 6 / W()
}
func W() float64 {
	return math.Sqrt(k / M())
}
func M() float64 {
	return p * v
}
