package main

import (
	"fmt"
	"math"
)

func maxInt(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxPriceRod(p []int, n int) int {

	fmt.Printf("p: %v, n: %v\n", p, n)

	if n == 0 {
		return 0
	}

	q := -math.MaxInt32

	for i := 0; i < n; i++ {
		q = maxInt(q, p[i]+maxPriceRod(p, n-i-1))
	}

	return int(q)
}

func main() {

	p := []int{1, 5, 8, 9} //, 10, 17, 17, 20, 24, 30}
	fmt.Printf("max cut price: %v\n", maxPriceRod(p, len(p)))

}
