package main

import "fmt"

func main() {
	var head *ListNode
	for i := 1; i <= 5; i++ {
		head = insert(head, i)
	}
	Show(reverseList(head))
}

func Show(l *ListNode) {
	p := l
	for p != nil {
		fmt.Printf("-> %v ", p.Val)
		p = p.Next
	}
}

func insert(root *ListNode, elem int) *ListNode {
	temp := ListNode{Val: elem, Next: nil}
	if root == nil {
		root = &temp
		return root
	}
	curr := root
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &temp
	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current, next := head, head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
