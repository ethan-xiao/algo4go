package main

import "sort"

func equalsIntArray(value []int, target []int) bool {
	if target == nil && value != nil {
		return false
	}
	if value == nil && target != nil {
		return false
	}
	size := len(value)
	if size != len(target) {
		return false
	}
	sort.Ints(value)
	sort.Ints(target)

	for i := 0; i < size; i++ {
		if value[i] != target[i] {
			return false
		}
	}
	return true
}
