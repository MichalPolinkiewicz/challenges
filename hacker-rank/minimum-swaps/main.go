package main

import "fmt"

//https://www.hackerrank.com/challenges/minimum-swaps-2/problem
func main() {
	a := []int32{7, 1, 3, 2, 4, 5, 6}
	fmt.Println(minimumSwaps(a))
}

func minimumSwaps(arr []int32) int32 {
	var counter int32
	for i, e := range arr {
		expected := i + 1
		if int(e) != expected {
			expectedElementId := findExpectedElementIndex(arr[i:], expected) + i
			arr[i], arr[expectedElementId] = arr[expectedElementId], arr[i]
			counter++
		}
	}
	return counter
}

func findExpectedElementIndex(arr []int32, x int) int {
	for i, e := range arr {
		if int(e) == x {
			return i
		}
	}
	return -1
}
