package main

import (
	"testing"
)

func TestTwoSumUnorderedList(t *testing.T) {
	value := twoSumUnorderedList([]int{2, 7, 11, 15}, 9)
	if !equalsIntArray(value, []int{0, 1}) {
		t.Fatal()
	}
	value = twoSumUnorderedList([]int{3, 2, 4}, 6)
	if !equalsIntArray(value, []int{1, 2}) {
		t.Fatal()
	}
}

func twoSumUnorderedList(nums []int, target int) []int {
	l := len(nums)
	//m := make(map[int]int) // RAM ~3.6MB
	m := make(map[int]int, l) // RAM ~2.9MB
	i, j, e := 0, 0, false
	for i = 0; i < l; i++ {
		if j, e = m[target-nums[i]]; e {
			return []int{i, j}
		}
		m[nums[i]] = i
	}
	return nil
}
