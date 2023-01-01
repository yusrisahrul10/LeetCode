package main

import "fmt"

func main() {
	var head ListNode = ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	Show(middleNode(&head))
}

func Show(l *ListNode) {
	p := l
	for p != nil {
		fmt.Printf("-> %v ", p.Val)
		p = p.Next
	}
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
