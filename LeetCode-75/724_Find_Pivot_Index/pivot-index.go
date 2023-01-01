package main

import "fmt"

func main() {
	values := []int{2, 1, -1}
	fmt.Println(pivotIndex(values))
}

func pivotIndex(nums []int) int {
	totalSum := 0
	leftSum := 0
	pivotIndex := -1
	for _, value := range nums {
		totalSum += value
	}

	for i, value := range nums {
		if totalSum-leftSum-value == leftSum {
			pivotIndex = i
			break
		} else {
			leftSum += value
		}
	}
	return pivotIndex
}
