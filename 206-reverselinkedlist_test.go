/**
leetcode #206
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
*/

package main

import (
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	if !equalsLinedList(reverseList(nil), nil) {
		t.Fatal()
	}
	if !equalsLinedList(reverseList(linerArray2LinkedList([]int{1})), linerArray2LinkedList([]int{1})) {
		t.Fatal()
	}
	if !equalsLinedList(reverseList(linerArray2LinkedList([]int{1, 2, 3, 4, 5})), linerArray2LinkedList([]int{5, 4, 3, 2, 1})) {
		t.Fatal()
	}
	if equalsLinedList(reverseList(linerArray2LinkedList([]int{1, 2, 3, 5, 4})), linerArray2LinkedList([]int{5, 4, 3, 2, 1})) {
		t.Fatal()
	}
}

func reverseList(head *ListNode) *ListNode {
	/*
		var n, p *ListNode
		for head != nil {
			n = p
			p = head
			head = head.Next
			p.Next = n
		}
	*/

	var h, n *ListNode = head, nil
	for h != nil {
		h, h.Next, n = h.Next, n, h
	}
	return n
}
