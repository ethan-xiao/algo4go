package main

import (
	"sort"
)

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

/**
Find Parent: x=i-1; (x-x%2)/2
*/
func linerArray2SymmetricTree(array []int) *TreeNode {
	size := len(array)
	if size == 0 {
		return nil
	}
	nodes := make([]*TreeNode, 0, size)
	root := &TreeNode{Val: array[0]}
	nodes = append(nodes, root)
	var p *TreeNode
	x := 0
	for i := 1; i < size; i++ {
		if array[i] == -1 {
			continue
		}
		x = i - 1
		p = nodes[int((x-x%2)/2)]
		if i%2 == 0 {
			p.Right = &TreeNode{Val: array[i]}
			nodes = append(nodes, p.Right)
		} else {
			p.Left = &TreeNode{Val: array[i]}
			nodes = append(nodes, p.Left)
		}
	}
	return root
}
