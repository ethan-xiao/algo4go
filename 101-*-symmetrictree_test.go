/**
leetcode #101

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following [1,2,2,null,3,null,3] is not:
    1
   / \
  2   2
   \   \
   3    3
*/

package main

import (
	"testing"
)

type TreeNode struct {
	Val   int       `json:"val,omitempty"`
	Left  *TreeNode `json:"left,omitempty"`
	Right *TreeNode `json:"right,omitempty"`
}

func TestSymmetricTree(t *testing.T) {
	if !isSymmetricTree(linerArray2SymmetricTree([]int{1, 2, 2, 3, 4, 4, 3})) {
		t.Fatal()
	}
	if isSymmetricTree(linerArray2SymmetricTree([]int{1, 2, 2, -1, 3, -1, 3})) {
		t.Fatal()
	}
}

func isSymmetricTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !compareSymmetricTreeNodeValue(root.Left, root.Right) {
		return false
	}
	if root.Left == nil || root.Right == nil {
		return true
	}
	return compareSymmetricTreeNode(root.Left, root.Right)
}

// for example [1,2,2,null,3,null,3], this function is called once
func compareSymmetricTreeNode(left, right *TreeNode) bool {
	if !compareSymmetricTreeNodeValue(left.Right, right.Left) {
		return false
	}
	if !compareSymmetricTreeNodeValue(left.Left, right.Right) {
		return false
	}
	if left.Right != nil && right.Left != nil {
		if !compareSymmetricTreeNode(left.Right, right.Left) {
			return false
		}
	}
	if left.Left != nil && right.Right != nil {
		if !compareSymmetricTreeNode(left.Left, right.Right) {
			return false
		}
	}
	return true
}

func compareSymmetricTreeNodeValue(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val
}

/*
func isSymmetric(root *TreeNode) bool {
    if root == nil{
        return true
    }
    if root.Left == nil && root.Right==nil{
        return true
    }
    return isSame(root.Left, root.Right)

}

// for example [1,2,2,null,3,null,3], this function is called twice
func isSame(p, q *TreeNode)bool{
    if p== nil && q==nil{
        return true
    }
    if q!=nil && p!= nil && p.Val == q.Val{
        return isSame(p.Left, q.Right) && isSame(p.Right, q.Left)
    }else{
        return false
    }
}
*/
