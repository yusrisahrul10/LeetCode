package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	var answer []string

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			answer = append(answer, "FizzBuzz")
		case i%3 == 0:
			answer = append(answer, "Fizz") // % = modulo
		case i%5 == 0:
			answer = append(answer, "Buzz")
		default:
			answer = append(answer, strconv.Itoa(i))
		}
	}

	return answer
}
