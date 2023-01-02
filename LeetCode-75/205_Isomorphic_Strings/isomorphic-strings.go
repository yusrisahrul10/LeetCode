package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
}

func isIsomorphic(s string, t string) bool {
	checked := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if checked[s[i]] {
			continue
		}
		checked[s[i]] = true

		for j := i + 1; j < len(s); j++ {
			fmt.Println("s[j]", s[j], "s[i]", s[i], "t[j]", t[j], "t[i]", t[i])
			if s[j] != s[i] && t[j] != t[i] {
				continue
			}
			if s[j] == s[i] && t[j] == t[i] {
				continue
			}
			return false
		}
	}
	return true
}

// func isIsomorphic(s string, t string) bool {
// 	var isSame bool
// 	stringIndexSame := make([]bool, len(s))
// 	for i := 1; i <= len(s); i++ {
// 		if string(s[i]) == string(s[i-1]) {
// 			stringIndexSame[i] = true
// 			stringIndexSame[i] = true
// 		} else {

// 		}
// 		// if string(s[i]) == "e" {
// 		// 	isSame = true
// 		// 	stringIndexSame[i] = true
// 		// } else {
// 		// 	isSame = false
// 		// 	stringIndexSame[i] = false
// 		// }
// 		fmt.Println("Index", i-1, "char", string(s[i-1]))
// 	}
// 	// fmt.Println(stringIndexSame)

// 	// for pos, char := range s {
// 	// 	isSame := [...]bool{false}
// 	// 	fmt.Println("Index", pos, "char", string(char), isSame)
// 	// }
// 	return isSame
// }
