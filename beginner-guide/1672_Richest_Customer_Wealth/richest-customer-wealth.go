package main

import "fmt"

func main() {
	var values = [][]int{
		{1, 5},
		{7, 3},
		{3, 99},
	}
	fmt.Println(maximumWealth(values))
}

func maximumWealth(accounts [][]int) int {
	var richest int
	var value int
	for i := 0; i < len(accounts); i++ {
		value = 0
		for j := 0; j < len(accounts[i]); j++ {
			value += accounts[i][j]
		}
		if value > richest {
			richest = value
		}
	}
	return richest
}
