/**
leetcode #125

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
*/

package main

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	if !isPalindrome(".,") {
		t.Fatal()
	}
	if !isPalindrome("a.") {
		t.Fatal()
	}
	if !isPalindrome("ab a") {
		t.Fatal()
	}
	if !isPalindrome("aba") {
		t.Fatal()
	}
	if !isPalindrome("aa") {
		t.Fatal()
	}
	if !isPalindrome("A man, a plan, a canal: Panama") {
		t.Fatal()
	}
	if !isPalindrome("Aman, a plan, a canal: Panama") {
		t.Fatal()
	}

	if isPalindrome("ab") {
		t.Fatal()
	}
	if isPalindrome("A man, a plan, a canal: Panaa") {
		t.Fatal()
	}
	if isPalindrome("Aman, a plan, a canal: Panaa") {
		t.Fatal()
	}
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	var a, b byte
	for i < j {
		for a = 0; i <= j; i++ {
			if a = toAlphanumeric(s[i]); a != 0 {
				break
			}
		}
		for b = 0; i <= j; j-- {
			if b = toAlphanumeric(s[j]); b != 0 {
				break
			}
		}
		if a != b {
			return false
		}
		i++
		j--
	}
	// i >= j
	// a == b or (a != b and (a == 0 or b == 0))
	return a == b
}

func toAlphanumeric(c byte) byte {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return c
	}
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return 0
}
