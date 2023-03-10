package main

import "fmt"

func main() {
	var head *ListNode
	for i := 1; i <= 5; i++ {
		head = insert(head, i)
	}
	Show(middleNode(head))
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

func middleNode(head *ListNode) *ListNode {

	fast, slow := head, head

	for fast != nil {

		fast = fast.Next

		if fast != nil {
			fast = fast.Next

		} else {
			// fast has reached the end of linked list
			// slow is on the middle point now
			break
		}

		slow = slow.Next
	}

	return slow

}
