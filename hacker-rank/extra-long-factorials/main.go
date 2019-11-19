package main

import (
	"fmt"
	"math/big"
)

//https://www.hackerrank.com/challenges/extra-long-factorials/problem
func main() {
	extraLongFactorials(25)
}

func extraLongFactorials(n int32) {
	var result = big.NewInt(1)

	for {
		if n > 0 {
			result = result.Mul(result, big.NewInt(int64(n)))
			n--
		} else {
			break
		}
	}
	fmt.Println(result)
}
