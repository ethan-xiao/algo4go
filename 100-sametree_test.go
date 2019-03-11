/**
leetcode #100

Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:

Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true
Example 2:

Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false
Example 3:

Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false
*/

package main

import "testing"

func TestSameTree(t *testing.T) {
	if !isSameTree(linerArray2SymmetricTree([]int{1, 2, 3}), linerArray2SymmetricTree([]int{1, 2, 3})) {
		t.Fatal()
	}
	if isSameTree(linerArray2SymmetricTree([]int{1, 2, 3}), linerArray2SymmetricTree([]int{1, 2, 4})) {
		t.Fatal()
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}