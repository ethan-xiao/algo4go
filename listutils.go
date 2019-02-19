package main

func equalsLinedList(p, q *ListNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return equalsLinedList(p.Next, q.Next)
}

func linerArray2LinkedList(array []int) *ListNode {
	size := len(array)
	if size == 0 {
		return nil
	}
	head := &ListNode{Val: array[0]}
	pre := head
	for i := 1; i < size; i++ {
		pre.Next = &ListNode{Val: array[i]}
		pre = pre.Next
	}
	return head
}
