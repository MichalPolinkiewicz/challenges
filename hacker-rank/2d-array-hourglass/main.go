package main

// https://www.hackerrank.com/challenges/2d-array/problem

func main() {

}

func hourglassSum(arr [][]int32) int32 {
	var results []int32
	var sum int32 = 0

	for i:= 0; i+2 < len(arr); i++ {
		arr1 := arr[i]
		arr2 := arr[i+1]
		arr3 := arr[i+2]
		for j:= 0; j+2 < len(arr[j]); j++ {
			sum = arr1[j] + arr1[j+1] + arr1[j+2] + arr2[j+1] + arr3[j] + arr3[j+1] + arr3[j+2]
			results = append(results, sum)
		}
	}

	max := results[0]
	for _, e := range results {
		if e > max {
			max = e
		}
	}
	return max
}
