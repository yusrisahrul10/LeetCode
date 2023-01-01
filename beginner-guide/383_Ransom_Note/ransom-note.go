package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aab", "abba"))
}

// func canConstruct(ransomNote string, magazine string) bool {
// 	var newString string
// 	// for i := 0; i < len(ransomNote); i++ {
// 	// 	if newString == ransomNote {
// 	// 		break
// 	// 	}
// 	// 	for j := 0; j < len(magazine); j++ {
// 	// 		var e1 byte = ransomNote[i]
// 	// 		var e2 byte = magazine[j]
// 	// 		var eString1 string = string(e1)
// 	// 		var eString2 string = string(e2)
// 	// 		if eString1 == eString2 {
// 	// 			newString += eString1
// 	// 			break
// 	// 		}
// 	// 	}
// 	// }
// 	if newString == ransomNote {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[rune]int)
	for _, val := range magazine {
		charCount[val] += 1
	}

	for _, val := range ransomNote {
		if charCount[val] == 0 {
			return false
		}

		charCount[val] -= 1
	}

	return true
}
