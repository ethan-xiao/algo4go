/**
leetcode #240
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
Example:

Consider the following matrix:

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
Given target = 5, return true.

Given target = 20, return false.
*/

package main

import (
	"testing"
)

func TestSearch2DMatrix2(t *testing.T) {
	array := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if !search2DMatrix(array, 1) {
		t.Fatal()
	}
	if !search2DMatrix(array, 4) {
		t.Fatal()
	}
	if !search2DMatrix(array, 17) {
		t.Fatal()
	}
	if search2DMatrix(array, 20) {
		t.Fatal()
	}
}

func search2DMatrix(matrix [][]int, target int) bool {
	m := len(matrix) - 1
	if m < 0 {
		return false
	}
	c := len(matrix[0]) - 1
	if c < 0 {
		return false
	}
	for n := 0; m >= 0 && n <= c; {
		if matrix[m][n] == target {
			return true
		}
		if matrix[m][n] > target {
			m--
		} else {
			n++
		}
	}
	return false
}
