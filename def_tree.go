package main

type TreeNode struct {
	Val   int       `json:"val,omitempty"`
	Left  *TreeNode `json:"left,omitempty"`
	Right *TreeNode `json:"right,omitempty"`
}
