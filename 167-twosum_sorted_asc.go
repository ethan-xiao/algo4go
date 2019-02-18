/**
leetcode #167

Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.
Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/

package main

import "testing"

func TestTwoSumSortedAscList(t *testing.T) {
	value := twoSumSortedAscList([]int{1, 2, 3, 6, 8, 11}, 10)
	if !equalsIntArray(value, []int{2, 5}) {
		t.Fatal()
	}
}

func twoSumSortedAscList(numbers []int, target int) []int {
	j := len(numbers) - 1
	v := 0
	for i := 0; i < j; {
		v = numbers[i] + numbers[j]
		if v == target {
			return []int{i + 1, j + 1}
		}
		if v > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
