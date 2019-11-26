package main

import (
	"fmt"
	"strings"
)

//https://www.hackerrank.com/challenges/array-left-rotation/
func main() {

}

func rotateLeft(a []int32, d int32) {
	x := a[d:]
	y := a[:d]
	rotated := append(x, y...)
	result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(rotated)), " "), "[]")
	fmt.Println(result)
}
