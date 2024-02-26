package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {

	sort.Strings(strs)
	prefix := []string{}
	for digit, char := range strs[0] {
		var common = true
		for _, word := range strs {
			if string(char) != string(word[digit]) {
				common = false
				break
			}
		}

		if common {
			prefix = append(prefix, string(char))
		} else {
			break
		}
	}
	return strings.Join(prefix, "")

}
