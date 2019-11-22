package main

import "fmt"

//https://www.hackerrank.com/challenges/ctci-recursive-staircase
func main() {
	fmt.Println(stepPerms(7))
}

func stepPerms(n int32) int32 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 4
	} else if n == 4 {
		return 7
	} else if n == 5 {
		return 13
	}
	return stepPerms(n-1) + stepPerms(n-2) + stepPerms(n-3)
}
