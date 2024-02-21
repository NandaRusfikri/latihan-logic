package main

import "strconv"

func main() {
	isPalindrome(-10)
	isPalindrome(121)
	isPalindrome(100)
}

func isPalindrome(x int) bool {

	text := strconv.Itoa(x)

	for i := 0; i < len(text); i++ {
		if text[i] == '-' || text[(len(text)-1)-i] == '-' {
			return false
		}

		if text[i] != text[(len(text)-1)-i] {
			return false
		}

	}
	return true

}
