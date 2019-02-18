/**
leetcode #136
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
*/

package main

import "testing"

func TestSingleNumber(t *testing.T) {
	if singleNumber([]int{2, 2, 1}) != 1 {
		t.Fatal()
	}

	if singleNumber([]int{2, 2, 1}) == 2 {
		t.Fatal()
	}

	if singleNumber([]int{1, 9, 1, 4, 5, 5, 4}) != 1 {
		t.Fatal()
	}

	if singleNumber([]int{3, 68, 3, 7, 10, 68, 7}) != 10 {
		t.Fatal()
	}

	if singleNumber([]int{2, 2, 2}) == 2 {
		t.Fatal()
	}
}

func singleNumber(nums []int) int {
	s := len(nums)
	for i := 1; i < s; i++ {
		nums[0] = nums[0] ^ nums[i]
	}
	return nums[0]
}
