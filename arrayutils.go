package main

import (
	"math"
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
l == Level, start from 0
i == Index of Array
p == Parent of Node
Find Parent: 2^l+(i-(i-1)%2)/2-l-1
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
	for l, i, j, c := 0.0, 1, 0, 0; i < size; {
		c = int(math.Pow(2, l+1))
		for j = 0; j < c && i < size; i, j = i+1, j+1 {
			if array[i] == -1 {
				continue
			}
			p = nodes[int(math.Pow(2, l))+(i-(i-1)%2)/2-int(l)-1]
			if i%2 == 0 {
				p.Right = &TreeNode{Val: array[i]}
				nodes = append(nodes, p.Right)
			} else {
				p.Left = &TreeNode{Val: array[i]}
				nodes = append(nodes, p.Left)
			}
		}
		l++
	}
	return root
}
