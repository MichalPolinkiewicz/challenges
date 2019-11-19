package main

import (
	"fmt"
	"math"
)

//https://www.hackerrank.com/challenges/new-year-chaos/problem
func main() {
	minimumBribes([]int32{2, 1, 5, 3, 4})
}

func minimumBribes(q []int32) {
	var counter int

	for i := len(q) - 1; i >= 0; i-- {
		e := q[i]
		if e-int32(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		for j := math.Max(0, float64(q[i] - 2)); int(j) < i; j++{
			if q[int(j)] > q[i] {
				counter++
			}
		}
	}
	fmt.Println(counter)
}
