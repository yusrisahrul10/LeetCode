package main

import "fmt"

func main() {
	var values = []int{1, 2, 3, 4}
	fmt.Println(runningSum(values))
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
