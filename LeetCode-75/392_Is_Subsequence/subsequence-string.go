package main

import "fmt"

func main() {
	fmt.Println("E", isSubsequence("ab", "baab"))
}

func isSubsequence(s string, t string) bool {
	var subsequence string
	index := 0
	for i := 0; i < len(t); i++ {
		for j := index; j < len(s); j++ {
			if string(t[i]) == string(s[j]) && index == j {
				subsequence += string(t[i])
				index++
				break
			}
		}
	}
	if subsequence == s {
		return true
	}
	return false
}
