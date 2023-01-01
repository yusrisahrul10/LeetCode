package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(123))
}

func numberOfSteps(num int) int {
	counter := 0
	limit := num
	for i := 0; i < limit; i++ {
		if num == 0 {
			break
		}
		if num%2 == 0 {
			num /= 2
			counter++
		} else {
			num -= 1
			counter++
		}
	}

	return counter
}
